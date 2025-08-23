package configs

// Connects to MongoDB and sets a Stable API version

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Client instance
var DB *mongo.Client

func ConnectDB() *mongo.Client {
	// Garante que EnvConfig está inicializado
	if EnvConfig == nil {
		log.Fatal("EnvConfig não foi inicializado. Chame LoadConfig() no início da aplicação.")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(EnvConfig.Database.Uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ Banco de dados não foi inicializado. Você esqueceu de chamar ConnectDB()?")
	}

	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collectionName)
	return collection
}
