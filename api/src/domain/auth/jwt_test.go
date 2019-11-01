package auth

import (
	. "bitcoin-api-docker/api/src/customtypes"
	"regexp"
	"testing"
)

func TestGenerateJwt(t *testing.T) {
	account := Account{
		UserID:   "fooo",
		Password: "barr",
		Birth:    "1992-01-01",
		Name:     "foo",
	}

	token, err := generateJwtToken(account)
	jwtRegex := regexp.MustCompile(`^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`)
	validToken := jwtRegex.MatchString(token)

	if err != nil || validToken == false {
		t.Error("invalid token")
	}

	userId, err := authenticateJwtToken(token)

	if err != nil || userId != account.UserID {
		t.Error("invalid token")
	}
}

func TestIvalidJwt(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI0OTI2NzMsInN1YiI6ImRhbmllbC5tb3JlaXJhIn0.xs3EAKlHrDTSnL8FzYFrj2OPYOvJR8R7FG6zc9H65t0"

	_, err := authenticateJwtToken(token)

	if err == nil {
		t.Error("Token should be expired")
	}

	token = ""

	_, err = authenticateJwtToken(token)

	if err == nil {
		t.Error("Token should be invalid")
	}
}
