package core

import "github.com/Akanibekuly/tsarka-tz/internal/interfaces"

type St struct {
	lg    interfaces.Logger
	cache interfaces.Cache
	db    interfaces.Db

	User interfaces.User
	Hash interfaces.Hash
}

func New(lg interfaces.Logger, cache interfaces.Cache, db interfaces.Db) *St {
	core := &St{
		lg:    lg,
		cache: cache,
		db:    db,
	}

	core.User = NewUserService(db)
	core.Hash = NewHashService(lg, db)

	return core
}
