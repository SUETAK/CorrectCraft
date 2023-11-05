package usecase

import (
	"connectrpc.com/connect"
	"context"
	"server/domain/repository"
	learningv1 "server/interfaces/proto/learning/v1"
)

type LearningAPI struct {
	//	依存する層が増えたらここに追加
	Repository repository.UserRepository
}

func (l *LearningAPI) Answer(ctx context.Context, req *connect.Request[learningv1.AnswerRequest]) (*connect.Response[learningv1.AnswerResponse], error) {
	res := connect.NewResponse(&learningv1.AnswerResponse{Sentence: "test"})
	res.Header().Set("Learning-Version", "v1")
	return res, nil
}
