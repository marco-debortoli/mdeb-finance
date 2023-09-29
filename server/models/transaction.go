package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// The collection used for transactions
const collectionName = "transactions"

// This is the transaction struct which represents the transaction type in the database
type Transaction struct {
	Notes    string               `bson:"notes"`
	Category string               `bson:"category"`
	Date     primitive.DateTime   `bson:"date"`
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Amount   primitive.Decimal128 `bson:"amount"`
}

// Private function to get the collection that is used to store transactions
func getTransactionCollection(database *mongo.Database) *mongo.Collection {
	return database.Collection(collectionName)
}

// Create a Transaction
func CreateTransaction(ctx context.Context, database *mongo.Database, transaction *Transaction) error {
	collection := getTransactionCollection(database)
	_, err := collection.InsertOne(ctx, transaction)
	return err
}
