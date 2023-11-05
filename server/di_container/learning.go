package di

import (
	"server/domain/usecase"
)

// NewLearningAPI 依存する層のインスタンス作成
func NewLearningAPI() *usecase.LearningAPI {
	return &usecase.LearningAPI{}
}
