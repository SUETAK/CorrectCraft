package migrations

import (
	"context"
	"github.com/uptrace/bun"
	"server/domain/repository/models"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().Model((*models.User)(nil)).IfExists().Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
