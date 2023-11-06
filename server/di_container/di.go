package di

import (
	"github.com/uptrace/bun"
	"net/http"
	"server/di_container/repository"
	"server/interfaces/proto/learning/v1/learningv1connect"
)

func InitLearning(mux *http.ServeMux, db *bun.DB) {
	answerRepository := di_container.NewAnswerRepository(db)
	userRepository := di_container.NewUserRepository(db)
	learning := NewLearningAPI(
		userRepository,
		answerRepository,
	)
	path, handler := learningv1connect.NewLearningServiceHandler(learning)
	mux.Handle(path, handler)
}
