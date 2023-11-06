package usecase

import (
	"connectrpc.com/connect"
	"context"
	learningv1 "server/interfaces/proto/learning/v1"
	interfaces "server/interfaces/repository"
)

type LearningAPI struct {
	//	依存する層が増えたらここに追加
	UserRepository   interfaces.UserRepository
	AnswerRepository interfaces.AnswerRepository
}

func (l LearningAPI) Answer(ctx context.Context, c *connect.Request[learningv1.AnswerRequest]) (*connect.Response[learningv1.AnswerResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (l LearningAPI) CreateAnswer(ctx context.Context, req *connect.Request[learningv1.AnswerRequest]) (*connect.Response[learningv1.AnswerResponse], error) {

	err := l.AnswerRepository.CreateAnswer(ctx, req.Msg.Sentence, req.Msg.UserId)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&learningv1.AnswerResponse{Sentence: "test"})
	res.Header().Set("Learning-Version", "v1")
	return res, nil
}
