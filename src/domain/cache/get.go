package cache

import "os"

func Get(key string) string {
	os.Getenv(key)
}
