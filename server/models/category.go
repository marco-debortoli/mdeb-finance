package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Name of the transaction category collection
const categoryCollectionName = "categories"

// These are the valid Category types
type CategoryType string

const (
	TYPE_DEBIT    CategoryType = "debit"
	TYPE_CREDIT   CategoryType = "credit"
	TYPE_TRANSFER CategoryType = "transfer"
)

// This is the struct which represents a category
type TransactionCategory struct {
	Name   string             `bson:"name" json:"name"`
	Type   CategoryType       `bson:"type" json:"type"`
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Active bool               `bson:"active" json:"active"`
}

// Function to get the TransactionCategory collection
func getCategoryCollection(database *mongo.Database) *mongo.Collection {
	return database.Collection(categoryCollectionName)
}

// Function to create a TransactionCategory from a string name and type
func CreateTransactionCategory(ctx context.Context, database *mongo.Database, name string, catType CategoryType) (TransactionCategory, error) {
	collection := getCategoryCollection(database)

	category := TransactionCategory{
		Name:   name,
		Type:   catType,
		Active: true,
	}

	res, err := collection.InsertOne(ctx, category)
	if err == nil {
		category.ID = res.InsertedID.(primitive.ObjectID)
	}

	return category, err
}

// Functions to retrieve a lists of TransactionCategory objects from the database
func GetTransactionCategoryList(ctx context.Context, database *mongo.Database, filter bson.M) []TransactionCategory {
	collection := getCategoryCollection(database)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []TransactionCategory

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return results
}

func GetAllTransactionCategories(ctx context.Context, database *mongo.Database) []TransactionCategory {
	return GetTransactionCategoryList(ctx, database, bson.M{})
}

func GetActiveTransactionCategories(ctx context.Context, database *mongo.Database) []TransactionCategory {
	return GetTransactionCategoryList(ctx, database, bson.M{"active": true})
}

func GetInactiveTransactionCategories(ctx context.Context, database *mongo.Database) []TransactionCategory {
	return GetTransactionCategoryList(ctx, database, bson.M{"active": false})
}

// Function to retrieve a single TransactionCategory by ObjectID
func GetTransactionCategory(ctx context.Context, database *mongo.Database, id primitive.ObjectID) (TransactionCategory, error) {
	collection := getCategoryCollection(database)

	result := collection.FindOne(ctx, bson.M{"_id": id})

	var found TransactionCategory
	err := result.Decode(&found)

	return found, err
}

// Update the active status of a TransactionCategory
func UpdateTransactionCategoryActive(ctx context.Context, database *mongo.Database, id primitive.ObjectID, active bool) error {
	collection := getCategoryCollection(database)

	_, err := collection.UpdateByID(ctx, id, bson.M{"active": active})

	return err
}
