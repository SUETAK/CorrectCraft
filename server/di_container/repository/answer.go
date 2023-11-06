package di_container

import (
	"github.com/uptrace/bun"
	"server/domain/repository"
	interfaces "server/interfaces/repository"
)

func NewAnswerRepository(db *bun.DB) interfaces.AnswerRepository {
	return &repository.AnswerRepository{
		Db: db,
	}
}
