package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	/*
		Sarama es un libreria de cliente en go para
		trabajar con apache kafka versiones 0.8 y
		posteriores. Incluye una API de alto nivel
		para producir y consumir mensajes facilmente,
		y una API  de bajo nivel para controlar bytes
	*/
	"github.com/edpo1998/kafka/sarama"
)

func main() {

	fmt.Println("Starting synchronous Kafka producer...")
	time.Sleep(5 * time.Second)
	// Creacion del canal
	doneCh := make(chan struct{})

	// Esta funcion genera mensajes cada 5 min.
	go func() {
		// Key del mensaje
		msgCount := 0
		for {
			// Creamos un Objeto de prueba
			msgCount++
			game := new(Game)
			game.Game_id = int64(msgCount)
			game.Players = 10
			game.Gamer_Name = "Random"
			game.Winner = 2
			game.Queue = "Kafka"
			sendToKafka(game)
			time.Sleep(5 * time.Second)
		}
	}()

	<-doneCh // Final del canal
}

type Game struct {
	Game_id    int64  `json:"game_id"`
	Players    int64  `json:"players"`
	Gamer_Name string `json:"gamer_name"`
	Winner     int64  `json:"winner"`
	Queue      string `json:"queue"`
}

func (g Game) ToJSON() []byte {
	ToJSON, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func brokerAddr() string {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if len(brokerAddr) == 0 {
		brokerAddr = "localhost:9092"
	}
	return brokerAddr
}

func topic() string {
	topic := os.Getenv("TOPIC")
	if len(topic) == 0 {
		topic = "default-topic"
	}
	return topic
}

func sendToKafka(game *Game) {

	// Instancia hacia sarama para iniciar el proceso de configuracion
	config := sarama.NewConfig()

	// Configuracion de retorno en los canalas deben colocarse en true cuando es SyncProducer
	config.Producer.Return.Successes = true // Devuelve los mensajes exitosos al canal
	config.Producer.Return.Errors = true    // Devuelve los mensajes erroneos al canal

	// Nivel de confiabilidad, por defecto viene en WaitForAll
	config.Producer.RequiredAcks = sarama.WaitForAll

	// Reintentos que realizara para poder enviar el mensaje
	config.Producer.Retry.Max = 5

	// Intermediario (middleware) entre el emisor y receptor
	// en este caso sera Kafka
	brokers := []string{brokerAddr()}

	// Inicia un SyncPorducer utilizando el broker y la configuracion
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	// Si se produjo un error lanzamos error y limpiamos recursos
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	// Variable para la configuracion de nuestros mensajes
	topic := topic()

	// Preparamos el mensaje
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(game.ToJSON()),
	}

	// Enviamos el mensaje
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	// Mensaje enviado al broker
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
