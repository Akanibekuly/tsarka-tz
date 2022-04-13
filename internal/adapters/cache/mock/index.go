package mock

import (
	"sync"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
)

type St struct {
	lg    interfaces.Logger
	store map[string]int
	sync.Mutex
}

func (m *St) New(lg interfaces.Logger) *St {
	return &St{
		lg: lg,
	}
}

func (m *St) Get(key string) (int, bool, error) {
	m.Lock()
	defer m.Unlock()

	val, ok := m.store[key]

	return val, ok, nil
}

func (m *St) Set(key string, value int, expiration time.Duration) error {
	m.Lock()
	defer m.Unlock()
	m.store[key] = value
	return nil
}

func (m *St) Del(key string) error {
	m.Lock()
	defer m.Unlock()
	delete(m.store, key)
	return nil
}
