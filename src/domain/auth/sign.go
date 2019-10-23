package auth

import (
	"fmt"
	"os"

	. "bitcoin-api/src/customtypes"
	. "bitcoin-api/src/infrastructure/aws/dynamo"

	"github.com/aws/aws-lambda-go/events"
)

func Sign(user User) (string, error) {
  validUser := Validate(user)

  if validUser != nil {
    return nil, errors.new{validUser}
  }

	err := Put(user, os.Getenv("REGISTER_USERS"))

	if err != nil {
    // if err == somethin {
    //   return nil, errors.New{"User already register!"}
    // }
		return nil, err
	}

	return "Account Created!", nil
}
