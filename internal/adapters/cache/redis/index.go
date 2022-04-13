package redis

import (
	"context"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/go-redis/redis/v8"
)

const requestTimeout = time.Second

type St struct {
	lg interfaces.Logger
	r  *redis.Client
}

func New(lg interfaces.Logger, url, psw string, db int) *St {
	return &St{
		lg: lg,
		r: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: psw,
			DB:       db,
		}),
	}
}

func (c *St) Get(key string) (int, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	data, err := c.r.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, false, nil
	}
	if err != nil {
		c.lg.Errorw("Redis: fail to 'get'", err)
		return 0, false, err
	}

	return data, true, nil
}

func (c *St) Set(key string, value int, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	err := c.r.Set(ctx, key, value, expiration).Err()
	if err != nil {
		c.lg.Errorw("Redis: fail to 'set'", err)
	}

	return err
}

func (c *St) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	err := c.r.Del(ctx, key).Err()
	if err != nil {
		c.lg.Errorw("Redis: fail to 'del'", err)
	}

	return err
}
