package customtypes

type Account struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Birth    string `json:"birth"`
}

type Transaction struct {
	UserId string
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Price  float32
	Date   string
}
