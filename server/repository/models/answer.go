package models

import (
	"github.com/uptrace/bun"
	"time"
)

// Answer answers テーブルのmodel ユーザーが回答した文と文法を保存する
type Answer struct {
	bun.BaseModel `bun:"table:answers,alias:a"`

	ID        string    `bun:"id,pk,type:uuid"`
	Sentence  string    `bun:"column:notnull,type:varchar(255)"`
	UserId    string    `bun:"column:notnull,type:uuid,column:user_id"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}
