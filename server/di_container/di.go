package di

import (
	"net/http"
	"server/interfaces/proto/learning/v1/learningv1connect"
)

func InitLearning(mux *http.ServeMux) {
	learning := NewLearningAPI()
	path, handler := learningv1connect.NewLearningServiceHandler(learning)
	mux.Handle(path, handler)
}
