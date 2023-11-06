package interfaces

import (
	"context"
	"server/domain/models"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) (*models.User, error)
	CreateUser(ctx context.Context, userName, password string) (*models.User, error)
}
