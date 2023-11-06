package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	UserName string    `bun:"type:varchar(255)"`
	Password string    `bun:"type:varchar(255)"`
}
