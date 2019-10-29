package customtypes

type Account struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Birth    string `json:"birth"`
}

type Transaction struct {
	UserId string
	Type   string `json:"type"`
	Amount string `json:"amount"`
	Price  float64
	Date   string
}

type SystemReport struct {
	UserId string `json:"userId"`
	Date   string `json:"date"`
}
