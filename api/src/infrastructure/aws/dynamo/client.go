package dynamo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func client() *dynamodb.DynamoDB {
	config := &aws.Config{Region: aws.String(os.Getenv("REGION"))}
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		config.Endpoint = aws.String("http://bitcoin-api-docker_dynamo_1:8000")
	}
	sess := session.Must(session.NewSession(config))

	dynamoClient := dynamodb.New(sess)

	return dynamoClient
}
