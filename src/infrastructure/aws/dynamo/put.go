package dynamo

import (
	"fmt"

	. "bitcoin-api/src/customtypes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Put(user User, TableName string) bool {
	dynamoClient := client()

	userMap, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		panic("Cannot marshal user into AttributeValue map")
		return false
	}

	params := &dynamodb.PutItemInput{
		TableName:           aws.String(TableName),
		Item:                userMap,
		ConditionExpression: aws.String("attribute_not_exists(userId)"),
		// ReturnValues:        aws.String("UPDATED_NEW"),
	}

	result, err := dynamoClient.PutItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return false
	}

	return true
}
