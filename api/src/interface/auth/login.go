package main

import (
	"encoding/json"

	. "bitcoin-api/src/customtypes"
	. "bitcoin-api/src/domain/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// type Response events.APIGatewayProxyResponse

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	json.Unmarshal([]byte(req.Body), &user)

	// Source prob will need to change
	source := req.RequestContext.Identity.SourceIP

	resp, err := Login(user, source)

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
