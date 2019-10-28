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

	// Source prob will need to change
	source := req.RequestContext.Identity.SourceIP

	resp, err := auth.Login(account, source)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(resp),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
