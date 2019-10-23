package customtypes

type Account struct {
	User  User
	Token string
}

type User struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}
