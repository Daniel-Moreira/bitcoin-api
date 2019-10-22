package json

import "encoding/json"

func toJson() {
	json.Unmarshal([]byte(req.Body), &user)
}

func toString() {
	stringU, _ := json.Marshal(user)
}
