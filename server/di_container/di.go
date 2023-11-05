package di

import (
	"net/http"
	"server/interfaces/proto/learning/v1/learningv1connect"
	"server/usecase"
)

func InitLearning(mux *http.ServeMux) {
	learning := usecase.NewLearningAPI()
	path, handler := learningv1connect.NewLearningServiceHandler(learning)
	mux.Handle(path, handler)
}
