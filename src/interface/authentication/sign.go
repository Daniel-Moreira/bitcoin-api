package main

import (
	"encoding/json"
	"fmt"

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

	// if !strings.Contains(user.UserID, "@") {
	// 	return Response{StatusCode: 500, Body: "Email address is required"}, nil
	// }
	// if len(user.Password) < 4 {
	// 	return Response{StatusCode: 500, Body: "Password must have at least 4 characters"}, nil
	// }
	// if user.Source == "" {
	// 	return Response{StatusCode: 500, Body: "Source is required"}, nil
	// }

	return Response{
		StatusCode: 200,
		Body:       "Hello"}, nil
}

func main() {
	lambda.Start(Handler)
}
