package auth

// import (
// 	jwt "github.com/dgrijalva/jwt-go"
// 	"fmt"
// )

// func generateJwtToken (user User) (string) {
// 	const jwtMethod = jwt.GetSigningMethod("HS256")
// 	const claims = &jwt.StandardClaims{
// 		Subject: user.UserID
// 		ExpiresAt: 15000,
//    	Issuer: user.source
// 	}

// 	const token = jwt.NewWithClaims(jwtMethod, claims)
// 	const signToken = token.SignedToken(user.password)

//   fmt.Println("token %v", signToken)

//   return signToken
// }

// func login (user User) (string, error) {
//   // match password
//   savedPass := dynamo.get(user.UserID)

//   if savedPass == nil {
//     return nil, errors.New("Unregistered user!")
//   }
//   else if savedPass != user.password {
//     return nil, errors.New("User passaword doesn't match!")
//   }

//   jwtToken := generateJwtToken(user)

//   return jwtToken, nil
// }
