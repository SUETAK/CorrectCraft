package models

import "github.com/uptrace/bun"

// SelectedGrammar selected_grammars テーブルのmodel　選択された文法を保存する
type SelectedGrammar struct {
	bun.BaseModel `bun:"table:selected_grammars,alias:sg"`

	ID       int64  `bun:"id,pk,autoincrement"`
	AnswerId int64  `bun:",notnull,column:answer_id"`
	Grammar  string `bun:",notnull,type:varchar(1024)"`

	Answer *Answer `bun:"rel:belongs-to,join:answer_id=id"`
}
