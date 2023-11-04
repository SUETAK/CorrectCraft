package models

import "github.com/uptrace/bun"

// SelectedGrammar selected_grammars テーブルのmodel　選択された文法を保存する
type SelectedGrammar struct {
	bun.BaseModel `bun:"table:selected_grammars,alias:sg"`

	ID       string `bun:"id,pk,type:uuid"`
	AnswerId string `bun:"type:uuid"`
	Grammar  string `bun:"notnull,type:varchar(255)"`

	Answer *Answer `bun:"rel:belongs-to,join:answer_id=id"`
}
