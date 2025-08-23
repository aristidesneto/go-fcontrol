package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func InitMongo() {
	log.Printf("Iniciando MongoDB...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().ApplyURI(EnvConfig.Database.Uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB não respondeu ao ping: %v", err)
	}

	log.Println("✅ Conectado ao MongoDB")

	MongoClient = client
	MongoDatabase = client.Database(EnvConfig.Database.Name)
}

func DisconnectMongo() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = MongoClient.Disconnect(ctx)
		log.Println("🛑 Conexão com MongoDB encerrada")
	}
}
