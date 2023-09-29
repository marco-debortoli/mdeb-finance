package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/marco-debortoli/mdeb-finance/server/models"
)

func main() {
	fmt.Printf("Connecting to MongoDB\n")

	cnxt := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017")

	client, err := mongo.Connect(cnxt, clientOptions)
	if err != nil {
		fmt.Printf("Could not connect to MongoDB")
		panic(err)
	}

	if err := client.Ping(cnxt, readpref.Primary()); err != nil {
		panic(err)
	}

	database := client.Database("finance")

	amount, _ := primitive.ParseDecimal128("100.00")

	transaction := models.Transaction{
		Date:     primitive.NewDateTimeFromTime(time.Now()),
		Notes:    "Testing Notes",
		Category: "Testing",
		Amount:   amount,
	}

	err = models.CreateTransaction(cnxt, database, &transaction)

	if err != nil {
		panic(err)
	}
}
