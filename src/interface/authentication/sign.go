package main

import (
	. "bitcoin-api/src/customtypes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	var user User
	json.Unmarshal([]byte(req.Body), &user)

	fmt.Println("user", user)
	// var body = ctx.body

	// if !strings.Contains(user.email, "@") Response{StatusCode: 500, Body: "Email address is required"}
	// if len(user.password) < 4 Response{StatusCode: 500, Body: "Password must have at least 4 characters"}
	// if user.source == nil Response{StatusCode: 500, Body: "Source is required"}

	return Response{StatusCode: 200, Body: "Hello"}, nil
}

func main() {
	lambda.Start(Handler)
}
