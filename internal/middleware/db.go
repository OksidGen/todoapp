package middleware

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var taskCollection *mongo.Collection

func init() {
	loadEnv()
	initializeDatabase()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file")
	}
}

func initializeDatabase() {
	connectionString := os.Getenv("DB_URI")

	dbName := os.Getenv("DB_NAME")

	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal("Failed to connect to cluster: ", err)
	}

	err = client.Ping(context, nil)
	if err != nil {
		log.Fatal("Failed to ping cluster: ", err)
	}

	log.Println("Connected to MongoDB!")

	database = client.Database(dbName)

	taskCollection = database.Collection(collName)
	log.Println("Connected to", database.Name(), "+", taskCollection.Name())
}
