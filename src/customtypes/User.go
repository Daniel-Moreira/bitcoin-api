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
