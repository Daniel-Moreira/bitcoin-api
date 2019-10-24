package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Put(TableName string, data interface{}) error {
	dynamoClient := client()

	userMap, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		panic("Cannot marshal data into AttributeValue map")
		return err
	}

	params := &dynamodb.PutItemInput{
		TableName:           aws.String(TableName),
		Item:                userMap,
		ConditionExpression: aws.String("attribute_not_exists(userId)"),
	}

	result, err := dynamoClient.PutItem(params)
	if err != nil {
		awsErr := err.(awserr.Error)

		if awsErr.Code() == "ConditionalCheckFailedException" {
			return errors.New("User already register")
		}

		return err
	}

	fmt.Println(result)

	return nil
}
