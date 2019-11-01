package main

import (
	"encoding/json"
	"os"
	"strings"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/auth"
	"bitcoin-api-docker/api/src/domain/trade"
	"bitcoin-api-docker/api/src/infrastructure/mysql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transaction Transaction
	json.Unmarshal([]byte(req.Body), &transaction)

	// Local token validation
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		token := strings.Split(req.Headers["Authorization"], " ")[1]
		userId, err := auth.Authenticate(token)

		transaction.UserId = userId
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       err.Error(),
			}, nil
		}
	} else {
		transaction.UserId = req.RequestContext.Authorizer["principalId"].(string)
	}

	mysql := &mysql.Mysql{}
	resp, err := trade.XChange(transaction, mysql)

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
