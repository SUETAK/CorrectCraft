package di

import (
	"github.com/uptrace/bun"
	"server/domain/repository"
	interfaces "server/interfaces/repository"
)

func NewRepository(db *bun.DB) interfaces.UserRepository {
	return repository.UserRepository{
		Db: db,
	}
}
