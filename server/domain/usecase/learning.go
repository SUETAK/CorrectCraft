package usecase

import (
	"connectrpc.com/connect"
	"context"
	learningv1 "server/interfaces/proto/learning/v1"
	interfaces "server/interfaces/repository"
)

type LearningAPI struct {
	//	依存する層が増えたらここに追加
	Repository interfaces.UserRepository
}

func (l *LearningAPI) Answer(ctx context.Context, req *connect.Request[learningv1.AnswerRequest]) (*connect.Response[learningv1.AnswerResponse], error) {
	res := connect.NewResponse(&learningv1.AnswerResponse{Sentence: "test"})
	res.Header().Set("Learning-Version", "v1")
	return res, nil
}
