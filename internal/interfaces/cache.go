package interfaces

import "time"

type Cache interface {
	Get(key string) (int, bool, error)
	Set(key string, value int, expiration time.Duration) error
	Del(key string) error
}
