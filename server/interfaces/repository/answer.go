package interfaces

import (
	"context"
)

type AnswerRepository interface {
	CreateAnswer(ctx context.Context, answerSentence string, userId string) error
}
