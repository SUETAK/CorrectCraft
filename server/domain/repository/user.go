package repository

import (
	"context"
	"github.com/uptrace/bun"
	"server/domain/models"
)

// UserRepository の依存するmodel が増えたらここに追加する？
type UserRepository struct {
	Db *bun.DB
}

func (ur UserRepository) GetAllUser(ctx context.Context) (*models.User, error) {
	user := new(models.User)
	err := ur.Db.NewSelect().Model(user).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}