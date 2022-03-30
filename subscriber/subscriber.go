package main

import (
	"fmt"
	"os"
	"os/signal"
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
	fmt.Println("Starting synchronous Kafka subscriber...")
	time.Sleep(5 * time.Second)

	// Instancia hacia sarama para iniciar el proceso de configuracion del Consumer
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.CommitInterval = 5 * time.Second
	config.Consumer.Return.Errors = true

	// Intermediario (middleware) entre el emisor y receptor
	// en este caso sera Kafka
	brokers := []string{brokerAddr()}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topic := topic()

	// decide about the offset here: literal value, sarama.OffsetOldest, sarama.OffsetNewest
	// this is important in case of reconnection
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {

			select {
			// Error
			case err := <-consumer.Errors():
				fmt.Println(err)
			// Se produjo un mensaje
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			// Se interrumpio el proceso del channel
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh                                       // Final del canal
	fmt.Println("Processed", msgCount, "messages") // Mensajes procesados
}

/*
	Funciones para verificacion de variables de
	entorno del broker y configuracion del topic
*/

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
