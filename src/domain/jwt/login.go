package login

import (
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

func login (user User) token string {
	const jwtMethod = jwt.GetSigningMethod("HS256")
	const claims = &jwt.StandardClaims{
		Subject: user.email
		ExpiresAt: 15000,
   		Issuer: user.source
	}

	const token = jwt.NewWithClaims(jwtMethod, claims)
	const signToken = token.SignedToken(user.password)

	fmt.printf("token %v", signToken)
}