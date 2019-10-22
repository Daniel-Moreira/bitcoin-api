package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
)

func sign (user User) () {
  if len(user.UserID) < 4 {
		return Response{StatusCode: 500, Body: "UserID must have at least 4 characters"}, nil
	}
	if len(user.Password) < 4 {
		return Response{StatusCode: 500, Body: "Password must have at least 4 characters"}, nil
	}
	if user.Source == "" {
		return Response{StatusCode: 500, Body: "Source is required"}, nil
	}
  
  addedUser := createAccount(user)
  
  if addedUser == nil {
    return Response{StatusCode: 200, Body: "User already registered!"}
  }
  
  return Response{StatusCode: 200, Body: "Account Created!"}
}