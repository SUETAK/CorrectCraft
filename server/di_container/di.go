package di

import (
	"connectrpc.com/connect"
	"github.com/uptrace/bun"
	"net/http"
	repositoryContainer "server/di_container/repository"
	"server/interfaces/proto/auth/v1/authv1connect"
	"server/jwt"
	contextkey "server/jwt/context_key"

	"server/interfaces/proto/learning/v1/learningv1connect"
)

func InitLearning(mux *http.ServeMux, db *bun.DB, key string, interceptor connect.Option) {
	answerRepository := repositoryContainer.NewAnswerRepository(db)
	userRepository := repositoryContainer.NewUserRepository(db)
	gptRepository := repositoryContainer.NewGptDescribeRepository(db)
	reader := contextkey.NewContextReader()
	openAi := NewOpenAIClient(key)
	learning := NewLearningAPI(
		userRepository,
		answerRepository,
		gptRepository,
		openAi,
		reader,
	)
	path, handler := learningv1connect.NewLearningServiceHandler(learning, interceptor)
	mux.Handle(path, handler)
}

func InitAuth(mux *http.ServeMux, db *bun.DB, issuer, keyPath string) {
	userRepository := repositoryContainer.NewUserRepository(db)
	tokenManager, err := jwt.NewTokenManager(issuer, keyPath)
	if err != nil {
		panic(err)
	}
	auth := NewAuthAPI(userRepository, tokenManager)
	path, handler := authv1connect.NewAuthServiceHandler(auth)
	mux.Handle(path, handler)
}
