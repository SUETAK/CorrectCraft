package di

import (
	"server/domain/usecase"
	"server/interfaces/proto/auth/v1/authv1connect"
	interfaces "server/interfaces/repository"
	"server/jwt"
)

func NewAuthAPI(repository interfaces.UserRepository, tokenManager jwt.ITokenManager) authv1connect.AuthServiceHandler {
	return &usecase.AuthAPI{
		UserRepository: repository,
		TokenManager:   tokenManager,
	}
}
