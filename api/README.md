<h1 id="bitcoin-api">Bitcoin API v1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Bitcoin API

## User login

<a id="opIdlogin"></a>

> Code samples

```javascript--go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/login", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /login`

Endpoint to user log into the system. It checks the user ID and password and returns a JWT token.

> Example responses

> 200 Response

```json
{
  "StatusCode": 200,
  "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzIzOTkxMDAsImlzcyI6IjEyNy4wLjAuMSIsInN1YiI6ImRhbmllbC5tb3JlaXJhIn0.igU_mYoe3CnYVbf-P6DyOL7yo0W8kCRiKyZKa2BFACU"
}
```

<h3 id="user-login-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|User successfully logged into the system|[JWTReponse](#schemajwtreponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Some error occur during the loging|[ErrorResponse](#schemaerrorresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## User sign up

<a id="opIdsign"></a>

> Code samples

```javascript--go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/sign", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /sign`

Endpoint to register an user into the system. It receives an user ID, password, full name and birthday.

> Example responses

> 201 Response

```json
{
  "StatusCode": 201,
  "Message": "Account Created!"
}
```

<h3 id="user-sign-up-endpoint-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|User successfully registered into the system|[SuccessResponse](#schemasuccessresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Some error occur during the sign up process|[ErrorResponse](#schemaerrorresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## Trade bitcoins

<a id="opIdxChangeBitcoin"></a>

> Code samples

```javascript--go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/xChangeBitcoin", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /xChangeBitcoin`

Endpoint for an authenticated user to trade bitcoins'. It receives the transaction type (buy or sell) and the amount of coins to trade.

> Example responses

> 201 Response

```json
{
  "StatusCode": 201,
  "Message": "Transaction registered!"
}
```

<h3 id="endpoint-for-an-authenticated-user-to-trade-bitcoins-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Transaction successfully registered into the system|[SuccessResponse](#schemasuccessresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Some error occur during the transaction process|[ErrorResponse](#schemaerrorresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## Reports

<a id="opIdreport"></a>

> Code samples

```javascript--go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/report", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /report`

Endpoint for an authenticated user to receive reports from the system'. Returns all transactions registered for an given user ID or day.

> Example responses

> 200 Response

```json
{
  "StatusCode": 200,
  "Report": [
        {
            "Amount": "0.000000",
            "Date": "2019-10-29 01:34:37",
            "Name": "Daniel Augusto de Melo Moreira",
            "Price": "0.000000",
            "TransactionId": "4",
            "Type": "buy",
            "UserId": "daniel.moreira"
        },
        {
            "Amount": "0.000000",
            "Date": "2019-10-29 01:36:42",
            "Name": "Daniel Augusto de Melo Moreira",
            "Price": "0.000000",
            "TransactionId": "5",
            "Type": "buy",
            "UserId": "daniel.moreira"
        }]
}
```

<h3 id="endpoint-for-an-authenticated-user-to-receive-reports-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Transaction successfully registered into the system|[ReportReponse](#schemareportreponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Some error occur during the report process|[ErrorResponse](#schemaerrorresponse)|

<aside class="success">
This operation does not require authentication
</aside>

