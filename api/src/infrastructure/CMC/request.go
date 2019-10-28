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

func GetBitcoinData() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)

	if err != nil {
		return '', err
	}

	q := url.Values{}
	q.Add("id", "1")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("BITCOIN_API_KEY"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return '', err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(respBody)

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	data := result["data"].(map[string]interface{})
	one := data["1"].(map[string]interface{})
	quote := one["quote"].(map[string]interface{})
	usd := quote["USD"].(map[string]interface{})

	return usd["price"], nil
}
