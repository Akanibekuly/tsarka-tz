package interfaces

import "github.com/Akanibekuly/tsarka-tz/internal/domain/entities"

type Hash interface {
	Calc(s []byte) (string, error)
	Result(id string) (*entities.HashSt, error)
}
