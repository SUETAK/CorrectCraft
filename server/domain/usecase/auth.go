package usecase

import (
	"connectrpc.com/connect"
	"context"
	v1 "server/interfaces/proto/auth/v1"
	interfaces "server/interfaces/repository"
)

type AuthAPI struct {
	UserRepository interfaces.UserRepository
}

func (a AuthAPI) CreateUser(ctx context.Context, c *connect.Request[v1.UserRequest]) (*connect.Response[v1.UserResponse], error) {
	user, err := a.UserRepository.CreateUser(ctx, c.Msg.UserName, c.Msg.Password)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.UserResponse{UserName: user.UserName}), nil
}
