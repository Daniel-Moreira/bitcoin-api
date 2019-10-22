package customtypes

type User struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
	Source   string `json:"sourceIp"`
}
