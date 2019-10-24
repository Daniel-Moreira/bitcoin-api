package cmc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	cache "src/domain/cache"
)

func GetBitcoinData() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)

	if err != nil {
		return err
	}

	q := url.Values{}
	q.Add("id", "1")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("BITCOIN_API_KEY"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	json.NewDecoder(resp.Body).Decode()
	respBody, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(respBody)

	now := time.Now()
	afterTime := now.Add(time.Minute * 15)

	cache.Put("BITCOIN_PRICE", "1")
	cache.Put("BITCOIN_EXP", afterTime.String())
}
