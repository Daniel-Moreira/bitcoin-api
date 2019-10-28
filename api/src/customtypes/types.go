package customtypes

type Account struct {
	User  User
	Name  string `json:"name"`
	Birth string `json:"birth"`
}

type User struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

type Transaction struct {
	UserId string
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Price  float32
	Date   string
}
