package repositoryContainer

import (
	"github.com/uptrace/bun"
	"server/domain/repository"
	interfaces "server/interfaces/repository"
)

func NewGptDescribeRepository(db *bun.DB) interfaces.GPTDescribeRepository {
	return &repository.GPTDescribeRepository{
		Db: db,
	}
}
