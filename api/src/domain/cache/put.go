package cache

import (
	"bitcoin-api-docker/api/src/infrastructure/aws/dynamo"
	"os"
)

func Put(key string, value string) error {
	return dynamo.Put(os.Getenv("CACHE_DB"), map[string]string{key: value})
}
