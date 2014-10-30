package logger

import (
	"time"

	redis "github.com/fzzy/radix/redis"
)

// Forwarder ...
type Forwarder struct {
	redisClient *redis.Client
	err         error
}

// NewForwarder : initalises a new Forwarder
func NewForwarder() *Forwarder {
	client, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)

	return &Forwarder{
		redisClient: client,
		err:         err,
	}
}

// Forward : sends the log on to the redis endpoint
func (f *Forwarder) Forward(log string) {
	if f.redisClient == nil {
		return
	}

	f.redisClient.Cmd("rpush", "logstash", log)
}
