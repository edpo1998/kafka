package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ADDRMONGO = "mongodb+srv://admin:1234@cluster0.fac6w.mongodb.net"
var NAMEDB = "User"
var NAMECOLL = "user"

type Game struct {
	Game_id    int64  `json:"game_id"`
	Players    int64  `json:"players"`
	Gamer_Name string `json:"gamer_name"`
	Winner     int64  `json:"winner"`
	Queue      string `json:"queue"`
}

func saveLogMongo(bjsonLog []byte) {

	// Get Struct Log
	logsobj := Game{}
	err := json.Unmarshal(bjsonLog, &logsobj)
	if err != nil {
		fmt.Println("Error al decodificar")
		return
	}

	// Database connection
	client, err := mongo.NewClient(options.Client().ApplyURI(ADDRMONGO))
	if err != nil {
		fmt.Println("Error de Conexion")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error de Conexion")
		return
	}
	defer client.Disconnect(ctx)

	dbLogs := client.Database(NAMEDB).Collection(NAMECOLL)

	// Insert Value
	_, err = dbLogs.InsertOne(ctx, logsobj)
	if err != nil {
		fmt.Println("Not inserted")
	} else {
		fmt.Println("New data inserted")
	}

}

func hellow() {

}
