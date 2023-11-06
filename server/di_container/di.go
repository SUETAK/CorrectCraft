package di

import (
	"github.com/uptrace/bun"
	"net/http"
	repositoryContainer "server/di_container/repository"

	"server/interfaces/proto/learning/v1/learningv1connect"
)

func InitLearning(mux *http.ServeMux, db *bun.DB) {
	answerRepository := repositoryContainer.NewAnswerRepository(db)
	userRepository := repositoryContainer.NewUserRepository(db)
	learning := NewLearningAPI(
		userRepository,
		answerRepository,
	)
	path, handler := learningv1connect.NewLearningServiceHandler(learning)
	mux.Handle(path, handler)
}
