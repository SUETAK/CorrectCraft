package di

import (
	"github.com/uptrace/bun"
	"net/http"
	"server/interfaces/proto/learning/v1/learningv1connect"
)

func InitLearning(mux *http.ServeMux, db *bun.DB) {
	repository := NewRepository(db)
	learning := NewLearningAPI(repository)
	path, handler := learningv1connect.NewLearningServiceHandler(learning)
	mux.Handle(path, handler)
}
