package di

import (
	"github.com/uptrace/bun"
	"net/http"
	repositoryContainer "server/di_container/repository"
	"server/interfaces/proto/auth/v1/authv1connect"

	"server/interfaces/proto/learning/v1/learningv1connect"
)

func InitLearning(mux *http.ServeMux, db *bun.DB, key string) {
	answerRepository := repositoryContainer.NewAnswerRepository(db)
	userRepository := repositoryContainer.NewUserRepository(db)
	openAi := NewOpenAIClient(key)
	learning := NewLearningAPI(
		userRepository,
		answerRepository,
		openAi,
	)
	path, handler := learningv1connect.NewLearningServiceHandler(learning)
	mux.Handle(path, handler)
}

func InitAuth(mux *http.ServeMux, db *bun.DB) {
	userRepository := repositoryContainer.NewUserRepository(db)
	auth := NewAuthAPI(userRepository)
	path, handler := authv1connect.NewAuthServiceHandler(auth)
	mux.Handle(path, handler)
}
