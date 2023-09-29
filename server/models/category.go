package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Name of the transaction category collection
const categoryCollectionName = "categories"

// This is the struct which represents a category
type TransactionCategory struct {
	Name     string             `bson:"name"`
	IsIncome bool               `bson:"is_income"`
	ID       primitive.ObjectID `bson:"_id"`
}

// Function to get the TransactionCategory collection
func getCategoryCollection(database *mongo.Database) *mongo.Collection {
	return database.Collection(categoryCollectionName)
}

// Function to create a TransactionCategory from a string and boolean
func CreateTransactionCategory(ctx context.Context, database *mongo.Database, name string, isIncome bool) error {
	collection := getCategoryCollection(database)

	category := TransactionCategory{
		Name:     name,
		IsIncome: isIncome,
	}

	_, err := collection.InsertOne(ctx, category)
	return err
}
