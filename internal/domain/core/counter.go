package core

import (
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
)

const (
	key           = "counter"
	expirationDur = time.Hour
)

func (c *St) Add(n int) (int, error) {
	val, err := c.Val()
	if err != nil && err != errs.ObjectNotFound {
		return 0, err
	}

	if err := c.cache.Set(key, val+n, expirationDur); err != nil {
		return 0, err
	}

	return val + n, nil
}

func (c *St) Sub(n int) (int, error) {
	val, err := c.Val()
	if err != nil && err != errs.ObjectNotFound {
		return 0, err
	}

	if err := c.cache.Set(key, val-n, expirationDur); err != nil {
		return 0, err
	}

	return val - n, nil
}

func (c *St) Val() (int, error) {
	val, ok, err := c.cache.Get(key)
	if err != nil {
		return val, err
	}
	if !ok {
		return val, errs.ObjectNotFound
	}

	return val, nil
}

func (c *St) Del() error {
	return c.cache.Del(key)
}
