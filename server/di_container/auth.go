package di

import (
	"server/domain/usecase"
	"server/interfaces/proto/auth/v1/authv1connect"
	interfaces "server/interfaces/repository"
)

func NewAuthAPI(repository interfaces.UserRepository) authv1connect.AuthServiceHandler {
	return &usecase.AuthAPI{
		UserRepository: repository,
	}
}
