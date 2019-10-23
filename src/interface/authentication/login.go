package main

import (
	"encoding/json"
	"fmt"

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
	user.Source = req.RequestContext.Identity.SourceIP

	stringU, _ := json.Marshal(user)

	fmt.Println(string(stringU))

	resp, err := Sign(user)

	return resp, err
}

func main() {
	lambda.Start(Handler)
}
