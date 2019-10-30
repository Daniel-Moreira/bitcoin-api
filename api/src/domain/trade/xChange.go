package trade

import (
	. "bitcoin-api-docker/api/src/customtypes"
	"bitcoin-api-docker/api/src/domain/cache"
	"bitcoin-api-docker/api/src/infrastructure/cmc"
	"bitcoin-api-docker/api/src/infrastructure/mysql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func XChange(transaction Transaction) (map[string]string, error) {
	now := time.Now()
	transaction.Date = now.Format("2006-01-02 15:04:05")

	lastBitcoinRequest := cache.Get("BITCOIN_EXP")
	timeFromLastRequest, _ := time.Parse("2006-01-02 15:04:05", lastBitcoinRequest)
	subDate := int((now.Sub(timeFromLastRequest)).Minutes())

	var bitCoinPrice float64
	bitCoinPrice, _ = strconv.ParseFloat(cache.Get("BITCOIN_PRICE"), 64)
	if subDate >= 0 {
		bitCoinPrice, err := cmc.GetBitcoinData()
		if err != nil {
			return nil, err
		}

		afterTime := now.Add(time.Minute * 15)

		cache.Put("BITCOIN_PRICE", fmt.Sprintf("%f", bitCoinPrice))
		cache.Put("BITCOIN_EXP", afterTime.Format("2006-01-02 15:04:05"))

		fmt.Println("Requesting")
	}

	transaction.Price = bitCoinPrice

	command := mysql.InsertCommand{
		TableName: os.Getenv("TRANSACTIONS_DB"),
	}
	command.Data = append(command.Data, transaction)

	err := mysql.Insert(command)

	if err != nil {
		return nil, err
	}

	return map[string]string{"Message": "Transaction registered!"}, nil
}
