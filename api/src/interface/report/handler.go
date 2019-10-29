package main

import (
	"encoding/json"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/auth"
	"bitcoin-api-docker/api/src/domain/trade"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SystemReport struct {
  UserId string,
  Date string,
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var report SystemReport
	json.Unmarshal([]byte(req.Body), &report)
	token := req.Headers["Jwt"]
	userId, err := auth.Authenticate(token)

	if report.UserId == nil && report.Date == nil {
    return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       string(map[string]string{"Error": "Report must have an userId or a date!"}),
		}, nil
  }

	resp, err := resport.New(report)

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
