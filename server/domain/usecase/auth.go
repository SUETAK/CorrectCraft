package usecase

import (
	"connectrpc.com/connect"
	"context"
	v1 "server/interfaces/proto/auth/v1"
	interfaces "server/interfaces/repository"
	"server/jwt"
	"time"
)

type AuthAPI struct {
	UserRepository interfaces.UserRepository
	TokenManager   jwt.ITokenManager
}

func (a AuthAPI) CreateUser(ctx context.Context, c *connect.Request[v1.UserRequest]) (*connect.Response[v1.UserResponse], error) {
	user, err := a.UserRepository.CreateUser(ctx, c.Msg.UserName, c.Msg.Password)
	if err != nil {
		return nil, err
	}

	token, err := a.TokenManager.CreateToken(user.ID.String(), 24*time.Hour)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&v1.UserResponse{UserName: user.UserName, Token: token}), nil
}

func (a AuthAPI) Login(ctx context.Context, c *connect.Request[v1.UserRequest]) (*connect.Response[v1.UserResponse], error) {
	user, err := a.UserRepository.GetUser(ctx, c.Msg.UserName)
	if err != nil {
		return nil, err
	}

	// TODO パスワードのハッシュ化
	if user.Password != c.Msg.Password {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	token, err := a.TokenManager.CreateToken(user.ID.String(), 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.UserResponse{UserName: user.UserName, Token: token}), nil
}
