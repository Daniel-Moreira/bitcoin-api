package cache

import "os"

func Put(key string, value string) error {
	return os.Setenv(key, value)
}
