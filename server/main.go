package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/marco-debortoli/mdeb-ledger/server/api"
	"github.com/marco-debortoli/mdeb-ledger/server/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Get the settings for our server
	godotenv.Load(".env")

	// Get the URL to connect to the database
	// This should be the full URL in format mongodb://<username>:<password>@<ip>:<port>
	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	// Get the port to run our server on
	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		log.Fatal("PORT environment variable not set")
	}

	// Connect to our database
	log.Println("Connecting to MongoDB database")

	clientOptions := options.Client()
	clientOptions.ApplyURI(dbUrl)

	// Set the default timeout for all operations to be 5 seconds
	clientOptions.SetTimeout(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the MongoDB database
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Select the appropriate "database" in MongoDB
	database := client.Database("ledger")

	apiConfig := api.APIConfig{
		DB:   database,
		Port: serverPort,
	}

	server.StartServer(&apiConfig)
}
