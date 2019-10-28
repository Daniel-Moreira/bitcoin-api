# bitcoin-api

# Install:
  1. Docker and docker-compose
  2. Run `docker-compose up`

## Test
  1. Make a HTTP POST Request to `localhost:3000/sign` with the following json:
      ```
      {
	      "userId":   "yourUserId",
	      "password": "yourPasswordId",
       	"name":     "yourFullName",
	      "birth":    "yourBirthdayDate" (ex.: "1992-12-01")
      }
      ```
  2. Make a HTTP GET Request to `localhost:3000/login` with the following json:
      ```
      {
	      "userId":   "yourUserId",
	      "password": "yourPasswordId"
      }
      ```
  3. Get the token returned from previous request and use it for all next requests adding to the HEADER
  4. If you want to eXChange bitcoins, make a HTTP POST Request to `localhost:3000/xChangeHandler` with the following json:
      ```
      {
        "type":     "buyOrSell",  (ex.: "buy")
	      "amount":   "amountToBuy" (ex.: 2)
      }
      ```
