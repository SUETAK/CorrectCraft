package interfaces

import (
	"context"
	"server/domain/models"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) (*models.User, error)
}
