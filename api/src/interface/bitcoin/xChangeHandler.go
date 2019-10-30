package main

import (
	"encoding/json"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/auth"
	"bitcoin-api-docker/api/src/domain/trade"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transaction Transaction
	json.Unmarshal([]byte(req.Body), &transaction)
	token := req.Headers["Jwt"]

	userId, err := auth.Authenticate(token)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	transaction.UserId = userId
	resp, err := trade.XChange(transaction)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	jsonResp, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       string(jsonResp),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
