package models

import (
	"context"
	"fmt"

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

// Set Account value and current value
func SetAccountValue(
	database *mongo.Database,
	account *Account,
	value float64,
	vTime primitive.DateTime,
) error {
	collection := getAccountCollection(database)

	newValue := AccountValue{
		Value: value,
		Date:  vTime,
	}

	// Need to figure out whether we are updating an existing value or adding a new one
	// Check and see if there is an entry for this date already
	var existingValue *AccountValue

	for _, v := range account.Values {
		if v.Date == vTime {
			existingValue = v
		}
	}

	if existingValue != nil {
		// We have an existing value - so let's update that one
		_, err := collection.UpdateOne(
			context.Background(),
			bson.M{"values.date": vTime, "_id": account.ID},
			bson.M{
				"$set": bson.M{
					"values.$.value": value,
				},
			},
		)

		if err != nil {
			fmt.Printf("Error setting nested account value for date: %v", vTime)
			return err
		}

		// This will update our passed in Account
		existingValue.Value = value
	} else {
		// We have a new value that we are adding - so let's add it
		_, err := collection.UpdateOne(
			context.Background(),
			bson.M{"_id": account.ID},
			bson.M{
				"$push": bson.M{
					"values": newValue,
				},
			},
		)

		if err != nil {
			fmt.Printf("Error adding new nested account value for date: %v", vTime)
			return err
		}

		// Update the list of values for our passed in Account
		account.Values = append(account.Values, &newValue)
	}

	// Now we need to update the current value if it is the most recent one
	if account.CurrentValue == nil || vTime >= account.CurrentValue.Date {
		_, err := collection.UpdateOne(
			context.Background(),
			bson.M{"_id": account.ID},
			bson.M{
				"$set": bson.M{
					"current_value": newValue,
				},
			},
		)

		if err != nil {
			fmt.Printf("Error setting current date: %v", err)
			return err
		}

		account.CurrentValue = &newValue
	}

	return nil
}

// Update old account value

// Mark inactive
