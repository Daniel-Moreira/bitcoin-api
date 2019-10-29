package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Get(TableName string, key map[string]string) (map[string]string, error) {
	dynamoClient := client()

	keyMap, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		return nil, fmt.Errorf("Cannot marshal %s into AttributeValue map", key)
	}

	params := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key:       keyMap,
	}

	result, err := dynamoClient.GetItem(params)
	if err != nil {
		return nil, err
	}

	var resultData map[string]string
	err = dynamodbattribute.UnmarshalMap(result.Item, resultData)
	if err != nil {
		return nil, errors.New("Failed to unmarshal Record" + err.Error())
	}

	fmt.Println(result.Item)
	return resultData, nil
}
