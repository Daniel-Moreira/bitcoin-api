package main

import (
	"bitcoin-api-docker/api/src/infrastructure/mysql"
	"encoding/json"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var account Account
	json.Unmarshal([]byte(req.Body), &account)

	mysql := &mysql.Mysql{}
	resp, err := auth.Sign(account, mysql)

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
