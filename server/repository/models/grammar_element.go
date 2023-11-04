package models

import "github.com/uptrace/bun"

// GrammarElements grammar_elements テーブルのmodel 文法の要素を保存する
type GrammarElements struct {
	bun.BaseModel `bun:"table:grammar_elements,alias:ge"`

	ID       string `bun:"id,pk,type:uuid"`
	AnswerId string `bun:"type:uuid"`
	Grammar  string `bun:"notnull,type:varchar(255)"`

	Answer *Answer `bun:"rel:belongs-to,join:answer_id=id"`
}
