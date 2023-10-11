package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/marco-debortoli/mdeb-finance/server/models"
)

func main() {
	fmt.Printf("Connecting to MongoDB\n")

	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://root:password@localhost:27017")

	// Set the default timeout for all operations to be 5 seconds
	clientOptions.SetTimeout(30 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Could not connect to MongoDB")
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	database := client.Database("ledger")

	// amount, _ := primitive.ParseDecimal128("100.00")

	/*transaction := models.Transaction{
		Date:     primitive.NewDateTimeFromTime(time.Now()),
		Notes:    "Testing Notes",
		Category: "Testing",
		Amount:   amount,
	}*/

	newCategory, err := models.CreateTransactionCategory(database, "Testing", "credit")
	fmt.Println(newCategory)
	if err != nil {
		panic(err)
	}

	foundResult, _ := models.GetTransactionCategory(database, newCategory.ID)

	resJson, _ := json.Marshal(foundResult)
	fmt.Println(string(resJson))

	/*
		activeResults := models.GetActiveTransactionCategories(ctx, database)
		inactiveResults := models.GetInactiveTransactionCategories(ctx, database)

		fmt.Println("Active Results")
		for _, result := range activeResults {
			resJson, _ := json.Marshal(result)
			fmt.Println(string(resJson))
		}

		fmt.Println("Inactive Results")
		for _, result := range inactiveResults {
			resJson, _ := json.Marshal(result)
			fmt.Println(string(resJson))
		}
	*/
}
