package main

import (
	"encoding/json"
	"fmt"

	. "bitcoin-api/src/domain"
	. "bitcoin-api/src/customtypes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	var user User
	json.Unmarshal([]byte(req.Body), &user)

	// Source prob will need to change
	user.Source = req.RequestContext.Identity.SourceIP

	stringU, _ := json.Marshal(user)

	fmt.Println(string(stringU))

	sign(user)

	return Response{
		StatusCode: 200,
		Body:       "Hello"}, nil
}

func main() {
	lambda.Start(Handler)
}
