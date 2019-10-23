package helpers

type Response events.APIGatewayProxyResponse

func BuildResponse (StatusCode int, payload interface{}) Response {
  if StatusCode == nil || payload == nil {
    return nil
  }

	return Response{
		StatusCode: StatusCode,
		Body: ToString(payload)
	}
}

func BuildResponse (StatusCode int, msg string) Response {
  if StatusCode == nil || msg == nil {
    return nil
  }
  
	return Response{
		StatusCode: StatusCode,
		Body: map[string]string{
			"Message": msg
		}
	}
}