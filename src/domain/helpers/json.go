package helpers

import "encoding/json"

func ToJson(from interface{}, *to interface{}) err error {
	err := json.Unmarshal([]byte(from), &to)
	return err
}

func ToString(object interface{}) (jsonString string, err error) {
	jsonData, err := json.Marshal(object)

	return string(jsonData), err
}
