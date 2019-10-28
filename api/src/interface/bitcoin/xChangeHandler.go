package main

import (
	"encoding/json"

	. "bitcoin-api/src/customtypes"
	. "bitcoin-api/src/domain/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transaction Transaction
	json.Unmarshal([]byte(req.Body), &transaction)

	transaction.userId = authenticate(header)
	resp, err := XChange(transaction)

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
