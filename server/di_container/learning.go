package di

import (
	"server/domain/usecase"
	"server/interfaces/proto/learning/v1/learningv1connect"
	interfaces "server/interfaces/repository"
)

// NewLearningAPI 依存する層のインスタンス作成
func NewLearningAPI(userRepository interfaces.UserRepository, answerRepository interfaces.AnswerRepository) learningv1connect.LearningServiceHandler {
	return &usecase.LearningAPI{
		UserRepository:   userRepository,
		AnswerRepository: answerRepository,
	}
}
