package di

import (
	"server/domain/usecase"
	contextkey "server/interfaces/context_key"
	openAIInterface "server/interfaces/open_ai"
	"server/interfaces/proto/learning/v1/learningv1connect"
	interfaces "server/interfaces/repository"
)

// NewLearningAPI 依存する層のインスタンス作成
func NewLearningAPI(
	userRepository interfaces.UserRepository,
	answerRepository interfaces.AnswerRepository,
	gptRepository interfaces.GPTDescribeRepository,
	openAIClient openAIInterface.OpenAI,
	reader contextkey.IContextReader) learningv1connect.LearningServiceHandler {
	return &usecase.LearningAPI{
		UserRepository:        userRepository,
		AnswerRepository:      answerRepository,
		GPTDescribeRepository: gptRepository,
		OpenAIClient:          openAIClient,
		ContextReader:         reader,
	}
}
