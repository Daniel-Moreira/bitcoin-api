package auth

import (
	"fmt"
	"os"

	. "bitcoin-api/src/customtypes"
	. "bitcoin-api/src/infrastructure/aws/dynamo"

	"github.com/aws/aws-lambda-go/events"
)

func Sign(user User) (events.APIGatewayProxyResponse, error) {
	if len(user.UserID) < 4 {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "UserID must have at least 4 characters"}, nil
	}
	if len(user.Password) < 4 {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Password must have at least 4 characters"}, nil
	}
	// if user.Source == "" {
	// 	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Source is required"}, nil
	// }

	fmt.Println(os.Getenv("REGISTER_USERS"))
	addedUser := Put(user, os.Getenv("REGISTER_USERS"))

	if addedUser == false {
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: "User already registered!"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Account Created!"}, nil
}
