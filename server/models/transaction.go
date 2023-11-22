package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// The collection used for transactions
const collectionName = "transactions"

// This is the transaction struct which represents the transaction type in the database
type Transaction struct {
	Name     string              `bson:"name" json:"name"`
	Category TransactionCategory `bson:"category" json:"category"`
	Date     primitive.DateTime  `bson:"date" json:"date"`
	ID       primitive.ObjectID  `bson:"_id,omitempty" json:"_id"`
	Amount   float64             `bson:"amount" json:"amount"`
}

// Private function to get the collection that is used to store transactions
func getTransactionCollection(database *mongo.Database) *mongo.Collection {
	return database.Collection(collectionName)
}

// Create a Transaction
func CreateTransaction(
	database *mongo.Database,
	name string,
	date primitive.DateTime,
	amount float64,
	category TransactionCategory,
) (Transaction, error) {
	collection := getTransactionCollection(database)

	transaction := Transaction{
		Category: category,
		Date:     date,
		Amount:   amount,
		Name:     name,
	}

	res, err := collection.InsertOne(context.Background(), transaction)
	if err == nil {
		transaction.ID = res.InsertedID.(primitive.ObjectID)
	}

	return transaction, err
}

// Get a List of transactions

type TransactionFilter struct {
	StartDate time.Time
	EndDate   time.Time
}

func GetTransactionList(database *mongo.Database, filter *TransactionFilter) []Transaction {
	collection := getTransactionCollection(database)

	transactionFilter := bson.M{}

	if filter != nil {
		innerFilter := bson.M{}
		if !filter.StartDate.IsZero() {
			innerFilter["$gte"] = primitive.NewDateTimeFromTime(filter.StartDate)
		}

		if !filter.EndDate.IsZero() {
			innerFilter["$lte"] = primitive.NewDateTimeFromTime(filter.EndDate)
		}

		transactionFilter["date"] = innerFilter
	}

	cursor, err := collection.Find(context.Background(), transactionFilter)
	if err != nil {
		panic(err)
	}

	results := []Transaction{}

	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}

	return results
}

// Delete a Transaction by _id
func DeleteTransaction(database *mongo.Database, id primitive.ObjectID) error {
	collection := getTransactionCollection(database)
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("Transaction with provided id does not exist")
	}

	return nil
}

// Retrieve a single Transaction by _id
func GetTransaction(database *mongo.Database, id primitive.ObjectID) (Transaction, error) {
	collection := getTransactionCollection(database)

	result := collection.FindOne(context.Background(), bson.M{"_id": id})

	var found Transaction
	err := result.Decode(&found)

	return found, err
}

// Update transaction (replace)
func UpdateTransaction(
	database *mongo.Database,
	transaction *Transaction,
	name string,
	date primitive.DateTime,
	amount float64,
	category TransactionCategory,
) error {
	collection := getTransactionCollection(database)

	update := bson.M{
		"$set": bson.M{
			"name":   name,
			"date":   date,
			"amount": amount,
			"category": bson.M{
				"name": category.Name,
				"_id":  category.ID,
			},
		},
	}

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": transaction.ID},
		update,
	)

	if err != nil {
		fmt.Printf("Failed to update transaction: %v", err)
		return err
	}

	// We should now update our transaction
	transaction.Name = name
	transaction.Date = date
	transaction.Category = category
	transaction.Amount = amount

	return nil
}
