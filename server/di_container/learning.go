package di

import (
	"server/domain/repository"
	"server/domain/usecase"
	"server/interfaces/proto/learning/v1/learningv1connect"
)

// NewLearningAPI 依存する層のインスタンス作成
func NewLearningAPI(repository repository.UserRepository) learningv1connect.LearningServiceHandler {
	return &usecase.LearningAPI{
		Repository: repository,
	}
}
