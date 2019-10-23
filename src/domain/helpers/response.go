package helpers

type Response events.APIGatewayProxyResponse

func BuildResponse (StatusCode int, payload interface{}) Response {
	return Response{
		StatusCode: StatusCode,
		Body: ToString(payload)
	}
}

func BuildResponse (StatusCode int, msg string) Response {
	return Response{
		StatusCode: StatusCode,
		Body: map[string]string{
			"Message": msg
		}
	}
}