package login

import (
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

jwt.GetSigningMethod("HS256")