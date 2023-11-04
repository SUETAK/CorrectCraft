package models

import "github.com/uptrace/bun"

// GptDescribe gpt_describes テーブルのmodel　GPT-4が解説した文章を保存する
type GptDescribe struct {
	bun.BaseModel `bun:"table:gpt_describes,alias:gd"`

	ID        string `bun:"id,pk,type:uuid"`
	Describe  string `bun:"column:notnull,type:varchar(255)"`
	AnswerId  string `bun:"type:uuid,column:answer_id"`
	CreatedAt string `bun:",nullzero,notnull,default:current_timestamp"`

	Answer *Answer `bun:"rel:belongs-to,join:answer_id=id"`
}
