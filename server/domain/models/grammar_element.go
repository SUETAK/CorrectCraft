package models

import "github.com/uptrace/bun"

// GrammarElement grammar_elements テーブルのmodel 文法の要素を保存する
type GrammarElement struct {
	bun.BaseModel `bun:"table:grammar_elements,alias:ge"`

	ID       int64  `bun:"id,pk,autoincrement"`
	AnswerId int64  `bun:",notnull,column:answer_id"`
	Grammar  string `bun:",notnull,type:varchar(255)"`

	Answer *Answer `bun:"rel:belongs-to,join:answer_id=id"`
}
