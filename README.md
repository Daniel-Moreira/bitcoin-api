# bitcoin-api

## Prerequisites:
  To run this project locally you'll need to have installed into your machine:
  * [Docker](https://docs.docker.com/v17.09/engine/installation/)
  * [Docker-compose](https://docs.docker.com/compose/install/)
  * [SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
  * [dep](https://golang.github.io/dep/docs/installation.html)
  
  To deploy this project you'll need only to install:
  * [serverless](https://serverless.com/framework/docs/providers/aws/guide/installation/)


## Getting Started
  You can start this project by following the steps:
  1. Run `docker-compose up -d`, which will build and start the database containers.
  2. At api folder run `make offline`, which will build and start the container for the api.

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
  3. Get the token returned from previous request and use it for all next requests adding it to the Bearer authorization in your HEADER http request.
  4. If you want to eXChange bitcoins, make a HTTP POST Request to `localhost:3000/xChangeHandler` with the following json:
      ```
      {
        "type":     "buyOrSell",  (ex.: "buy")
	      "amount":   "amountToBuy" (ex.: 2)
      }
      ```
  5. If you want a report status, make a HTTP GET Request to `localhost:3000/report` with the following json:
      ```
      {
	      "userId": "yourUserId" (ex.: "daniel.moreira")
      }
      ```
      or
      ```
      {
	      "date":"yourDate" (ex.: "2009-10-30")
      }
      ```
