package core

import (
	"hash/crc64"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/google/uuid"
)

type HashService struct {
	lg interfaces.Logger
	db interfaces.Db
}

var partitionTable = crc64.MakeTable(crc64.ISO)

func (h *HashService) Calc(s string) (string, error) {
	id := uuid.New().String()

	if err := h.db.HashCreate(id); err != nil {
		return "", err
	}

	go h.routine(s, id)

	return id, nil
}

func (h *HashService) routine(s, id string) {
	sum := crc64.Checksum([]byte(s), partitionTable)
	start := time.Now()

	for time.Since(start) < time.Minute {
		sum &= uint64(time.Now().UnixNano())
		time.Sleep(time.Minute)
	}

	err := h.db.HashUpdate(id, &entities.HashSt{
		Status: "FINISHED",
		Result: countOnes(sum),
	})
	if err != nil {
		h.lg.Errorw("[HASH] routine", err, "id", id, "str", s)
	}
}

func (h *HashService) Result(id string) (*entities.HashSt, error) {
	return h.db.HashGet(id)
}

func countOnes(n uint64) int {
	res := 0
	for n > 0 {
		res += int(n & 1)
		n >>= 1
	}

	return res
}
