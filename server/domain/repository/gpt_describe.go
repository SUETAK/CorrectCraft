package repository

import (
	"context"
	"github.com/uptrace/bun"
	"server/domain/models"
)

type GPTDescribeRepository struct {
	Db *bun.DB
}

func (g GPTDescribeRepository) CreateGPTDescribe(ctx context.Context, answerId int64, describe string) error {
	_, err := g.Db.NewInsert().Model(&models.GptDescribe{AnswerId: answerId, Describe: describe}).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
