package cache

import "os"

func Get(key string) string {
	return os.Getenv(key)
}
