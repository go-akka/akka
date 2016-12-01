package actor

import (
	"time"
)

type Config interface {
	GetBoolean(key string) bool
	GetString(key string) string
	GetInt(key string) int
	GetMillisDuration(key string) time.Duration
	GetConfig(path string) Config
}
