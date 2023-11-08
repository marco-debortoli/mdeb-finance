package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// The collection used for transactions
const collectionName = "transactions"

// This is the transaction struct which represents the transaction type in the database
type Transaction struct {
	Name     string             `bson:"name" json:"name"`
	Category string             `bson:"category" json:"category"`
	Date     primitive.DateTime `bson:"date" json:"date"`
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Amount   float64            `bson:"amount" json:"amount"`
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
		Category: category.Name,
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
