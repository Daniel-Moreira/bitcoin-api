package main

import (
	"encoding/json"
	"os"
	"strings"

	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/auth"
	"bitcoin-api-docker/api/src/domain/reports"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var report SystemReport
	json.Unmarshal([]byte(req.Body), &report)

	// Local token validation
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		token := strings.Split(req.Headers["Authorization"], " ")[1]
		_, err := auth.Authenticate(token)

		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       err.Error(),
			}, nil
		}
	}

	resp, err := reports.New(report)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	jsonResp, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResp),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
