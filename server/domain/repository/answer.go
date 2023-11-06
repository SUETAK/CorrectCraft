package repository

import (
	"context"
	"github.com/uptrace/bun"
	"server/domain/models"
)

type AnswerRepository struct {
	Db *bun.DB
}

func (a AnswerRepository) CreateAnswer(ctx context.Context, answerSentence string) error {
	answer := &models.Answer{Sentence: answerSentence}
	_, err := a.Db.NewInsert().Model(answer).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
