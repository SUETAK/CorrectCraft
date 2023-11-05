package di

import (
	"server/domain/usecase"
	"server/interfaces/proto/learning/v1/learningv1connect"
	interfaces "server/interfaces/repository"
)

// NewLearningAPI 依存する層のインスタンス作成
func NewLearningAPI(repository interfaces.UserRepository) learningv1connect.LearningServiceHandler {
	return &usecase.LearningAPI{
		Repository: repository,
	}
}
