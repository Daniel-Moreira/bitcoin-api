package cache

import (
	"bitcoin-api-docker/api/src/infrastructure/aws/dynamo"
	"os"
)

func Get(key string) string {
	result, _ := dynamo.Get(os.Getenv("CACHE_DB"), map[string]string{"key": key})

	return string(result["value"])
}
