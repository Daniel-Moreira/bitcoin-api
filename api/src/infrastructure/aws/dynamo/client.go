package dynamo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func client() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("REGION")),
		Endpoint: aws.String("http://localhost:8000"),
	}))

	dynamoClient := dynamodb.New(sess)

	return dynamoClient
}
