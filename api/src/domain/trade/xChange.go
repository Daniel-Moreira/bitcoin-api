package auth

import (
	"bitcoin-api-docker/api/src/domain/cache"
	"fmt"
	"os"
	"time"

	. "bitcoin-api/src/customtypes"
	. "bitcoin-api/src/infrastructure/mysql"
)

func XChange(transaction Transaction) (map[string]string, error) {
	now := time.Now()
	transaction.Date = now

	fmt.Printf("%v".transaction)
	if err != nil {
		return nil, err
	}

	lastBitcoinRequest := cache.Get("BITCOIN_EXP")
	int((now.Sub(lastBitcoinRequest)).Minutes()
	now := time.Now()
	afterTime := now.Add(time.Minute * 15)

	cache.Put("BITCOIN_PRICE", usd["price"])
	cache.Put("BITCOIN_EXP", afterTime.String())

	err = Put(os.Getenv("REGISTER_USERS"), user)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Message": "Account Created!"}, nil
}
