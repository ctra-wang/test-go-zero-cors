package redistest

import (
	"time"

	"github.com/Holyson/test-go-zero-cors/core/lang"
	"github.com/Holyson/test-go-zero-cors/core/stores/redis"
	"github.com/alicebob/miniredis/v2"
)

// CreateRedis returns a in process redis.Redis.
func CreateRedis() (r *redis.Redis, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	return redis.New(mr.Addr()), func() {
		ch := make(chan lang.PlaceholderType)

		go func() {
			mr.Close()
			close(ch)
		}()

		select {
		case <-ch:
		case <-time.After(time.Second):
		}
	}, nil
}
