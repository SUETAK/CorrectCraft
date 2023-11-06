package repositoryContainer

import (
	"github.com/uptrace/bun"
	"server/domain/repository"
	interfaces "server/interfaces/repository"
)

func NewUserRepository(db *bun.DB) interfaces.UserRepository {
	return repository.UserRepository{
		Db: db,
	}
}
