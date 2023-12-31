package repository

import (
	"context"
	"github.com/uptrace/bun"
	"server/domain/models"
)

type AnswerRepository struct {
	Db *bun.DB
}

func (a AnswerRepository) CreateAnswer(ctx context.Context, answerSentence, userId string) (int64, error) {
	answer := &models.Answer{Sentence: answerSentence, UserId: userId}
	_, err := a.Db.NewInsert().Model(answer).Exec(ctx)
	if err != nil {
		return 0, err
	}
	return answer.ID, nil
}
