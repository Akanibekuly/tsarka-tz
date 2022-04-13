package core

import (
	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/google/uuid"
)

type HashService struct {
	lg interfaces.Logger
	db interfaces.Db
}

func (h *HashService) Calc(s string) (string, error) {
	id := uuid.New().String()

	if err := h.db.HashCreate(id); err != nil {
		return "", err
	}

	go h.routine(s, id)

	return id, nil
}

func (h *HashService) routine(s, id string) {

}

func (h *HashService) Result(id string) (*entities.HashSt, error) {
	return h.db.HashGet(id)
}
