package sign

import (
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

type User struct {
	email string `json:"email"`
	password string `json:"password"`
	source string `json:"source"`
}

func Handler (ctx context.Context) (Response, error) {
	var body = ctx.body
	var user = &User{body}

	if !strings.Contains(user.email, "@") Response{StatusCode: 500, Body: "Email address is required"}
	if len(user.password) < 4 Response{StatusCode: 500, Body: "Password must have at least 4 characters"}
	if user.source == nil Response{StatusCode: 500, Body: "Source is required"}
}