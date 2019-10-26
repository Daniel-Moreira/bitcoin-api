package helpers

import "github.com/aws/aws-lambda-go/events"

type Response events.APIGatewayProxyResponse

func BuildResponse(StatusCode int, payload string) Response {
	return Response{
		StatusCode: StatusCode,
		Body:       payload,
	}
}
