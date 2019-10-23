package dynamo

import (
	"fmt"

	. "bitcoin-api/src/customtypes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func get(user User, TableName string) *User {
	dynamoClient := client()

	keyMap, err := dynamodbattribute.MarshalMap(user.UserID)
	if err != nil {
		panic("Cannot marshal UserID into AttributeValue map")
	}

	params := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key:       keyMap,
	}

	result, err := dynamoClient.GetItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return nil
	}

	resultUser := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &resultUser)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		return nil
	}

	fmt.Println("Success")
	fmt.Println(resultUser)

	return &resultUser
}
