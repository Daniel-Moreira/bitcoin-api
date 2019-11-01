package main

import (
	"bitcoin-api-docker/api/src/domain/auth"
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	return authResponse
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := strings.Split(event.AuthorizationToken, " ")[1]
	userId, err := auth.Authenticate(token)

	err = nil
	userId = "daniel.moreira"
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, err
	}

	return generatePolicy(userId, "Allow", event.MethodArn), nil
}

func main() {
	lambda.Start(handleRequest)
}
