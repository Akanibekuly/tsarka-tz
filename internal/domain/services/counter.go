package services

import (
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
)

const (
	key           = "counter"
	expirationDur = time.Hour
)

type CounterSt struct {
	lg    interfaces.Logger
	cache interfaces.Cache
}

func NewCounterService(lg interfaces.Logger, cache interfaces.Cache) *CounterSt {
	return &CounterSt{
		lg:    lg,
		cache: cache,
	}
}

func (c *CounterSt) Add(n int) (int, error) {
	val, err := c.Val()
	if err != nil && err != errs.ObjectNotFound {

		return 0, err
	}

	if err := c.cache.Set(key, val+n, expirationDur); err != nil {
		return 0, err
	}

	return val + n, nil
}

func (c *CounterSt) Sub(n int) (int, error) {
	val, err := c.Val()
	if err != nil && err != errs.ObjectNotFound {
		return 0, err
	}

	if err := c.cache.Set(key, val-n, expirationDur); err != nil {
		return 0, err
	}

	return val - n, nil
}

func (c *CounterSt) Val() (int, error) {
	val, ok, err := c.cache.Get(key)
	if err != nil {
		return val, err
	}
	if !ok {
		return val, errs.ObjectNotFound
	}

	return val, nil
}

func (c *CounterSt) Del() error {
	return c.cache.Del(key)
}
