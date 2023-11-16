package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Name of the account collection
const accountCollectionName = "accounts"

// These are the valid Account types
type AccountType string

const (
	TYPE_ACCOUNT_DEBIT      AccountType = "debit"
	TYPE_ACCOUNT_CREDIT     AccountType = "credit"
	TYPE_ACCOUNT_INVESTMENT AccountType = "investment"
)

// The following structs represent an Account and the nested Account Value
type AccountValue struct {
	Value float64            `bson:"value" json:"value"`
	Date  primitive.DateTime `bson:"date" json:"date"`
}

type Account struct {
	Name           string              `bson:"name" json:"name"`
	Type           AccountType         `bson:"type" json:"type"`
	ID             primitive.ObjectID  `bson:"_id,omitempty" json:"_id"`
	Active         bool                `bson:"active" json:"active"`
	IncludeNetWork bool                `bson:"include_net_worth" json:"include_net_worth"`
	StartDate      primitive.DateTime  `bson:"start_date" json:"start_date"`
	EndDate        *primitive.DateTime `bson:"end_date" json:"end_date"`
	CurrentValue   *AccountValue       `bson:"current_value,omitempty" json:"current_value"`
	Values         []*AccountValue     `bson:"values" json:"values"`
}

// Function to get the Account collection
func getAccountCollection(database *mongo.Database) *mongo.Collection {
	return database.Collection(accountCollectionName)
}

// This function will create a new Account
func CreateAccount(
	database *mongo.Database,
	name string,
	acctType AccountType,
	includeNetWorth bool,
	startDate primitive.DateTime,
) (Account, error) {
	collection := getAccountCollection(database)

	account := Account{
		Name:           name,
		Type:           acctType,
		Active:         true,
		IncludeNetWork: includeNetWorth,
		StartDate:      startDate,
		CurrentValue:   nil,
		Values:         make([]*AccountValue, 0),
		EndDate:        nil,
	}

	res, err := collection.InsertOne(context.Background(), account)
	if err == nil {
		account.ID = res.InsertedID.(primitive.ObjectID)
	}

	return account, err
}

// Get a single account
func GetAccount(database *mongo.Database, id primitive.ObjectID) (Account, error) {
	collection := getAccountCollection(database)

	result := collection.FindOne(context.Background(), bson.M{"_id": id})

	var found Account
	err := result.Decode(&found)

	return found, err
}

// Set Account current value
func SetAccountCurrentValue(
	database *mongo.Database,
	account *Account,
	value float64,
	vTime time.Time,
) error {
	collection := getAccountCollection(database)

	newValue := AccountValue{
		Value: value,
		Date:  primitive.NewDateTimeFromTime(vTime),
	}

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": account.ID},
		bson.M{
			"$set": bson.M{
				"current_value": newValue,
			},
			"$push": bson.M{
				"values": newValue,
			},
		},
	)

	if err != nil {
		fmt.Printf("Error setting account current value: %v", err)
		return err
	}

	// Set our provided account current value
	account.CurrentValue = &newValue
	account.Values = append(account.Values, &newValue)

	return nil
}

// Update old account value

// Mark inactive
