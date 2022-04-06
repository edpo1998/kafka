package bases

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ADDRMONGO = "mongodb://admindb:1234@35.208.184.236:27017"
var NAMEDB = "fase2"
var NAMECOLL = "fase2"

type Game struct {
	Juegoid        string `json:"juegoid"`
	Cantjugadores  string `json:"cantjugadores"`
	Nombrejuego    string `json:"nombrejuego"`
	Jugadorganador int    `json:"jugadorganador"`
	Queue_rabbit   string `json:"queue_rabbit"`
}

func SaveLogMongo(bjsonLog []byte) {

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
		fmt.Println("Conection Failed Mongo Cluster")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Failure Start Connect Client")
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
