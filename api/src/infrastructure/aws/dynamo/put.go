package dynamo

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Put(TableName string, data interface{}) error {
	dynamoClient := client()

	dataMap, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return errors.New("Cannot marshal data into AttributeValue map")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      dataMap,
	}

	_, err = dynamoClient.PutItem(params)
	if err != nil {
		awsErr := err.(awserr.Error)

		if awsErr.Code() == "ConditionalCheckFailedException" {
			return errors.New("User already register")
		}

		return err
	}

	return nil
}
