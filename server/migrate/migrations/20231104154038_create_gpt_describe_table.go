package migrations

import (
	"context"
	"server/repository/models"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().Model((*models.GptDescribe)(nil)).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().Model((*models.GptDescribe)(nil)).IfExists().Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}