package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/marco-debortoli/mdeb-finance/server/models"
)

func main() {
	fmt.Printf("Connecting to MongoDB\n")

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017")

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

	newCategory, err := models.CreateTransactionCategory(ctx, database, "Testing", "credit")
	fmt.Println(newCategory)
	if err != nil {
		panic(err)
	}

	foundResult, _ := models.GetTransactionCategory(ctx, database, newCategory.ID)

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
