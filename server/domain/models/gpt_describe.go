package models

import "github.com/uptrace/bun"

// GptDescribe gpt_describes テーブルのmodel　GPT-4が解説した文章を保存する
type GptDescribe struct {
	bun.BaseModel `bun:"table:gpt_describes,alias:gd"`

	ID        int64  `bun:"id,pk,autoincrement"`
	Describe  string `bun:",notnull,type:varchar(1024)"`
	AnswerId  int64  `bun:",notnull,column:answer_id"`
	CreatedAt string `bun:",nullzero,notnull,default:current_timestamp"`

	Answer *Answer `bun:"rel:belongs-to,join:answer_id=id"`
}
