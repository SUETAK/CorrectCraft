package usecase

import (
	"connectrpc.com/connect"
	"context"
	contextkey "server/interfaces/context_key"
	openAIInterface "server/interfaces/open_ai"
	learningv1 "server/interfaces/proto/learning/v1"
	interfaces "server/interfaces/repository"
)

type LearningAPI struct {
	//	依存する層が増えたらここに追加
	ContextReader    contextkey.IContextReader
	UserRepository   interfaces.UserRepository
	AnswerRepository interfaces.AnswerRepository
	OpenAIClient     openAIInterface.OpenAI
}

func (l LearningAPI) Answer(ctx context.Context, c *connect.Request[learningv1.AnswerRequest]) (*connect.Response[learningv1.AnswerResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (l LearningAPI) CreateAnswer(ctx context.Context, req *connect.Request[learningv1.AnswerRequest]) (*connect.Response[learningv1.AnswerResponse], error) {
	var uid string
	var err error
	if uid, err = l.ContextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	err = l.AnswerRepository.CreateAnswer(ctx, req.Msg.Sentence, uid)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&learningv1.AnswerResponse{Sentence: "test"})
	res.Header().Set("Learning-Version", "v1")
	return res, nil
}
