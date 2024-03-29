package main

import (
	"encoding/json"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// type Response events.APIGatewayProxyResponse

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var account Account
	json.Unmarshal([]byte(req.Body), &account)

	resp, err := auth.Login(account)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	jsonResp, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResp),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
