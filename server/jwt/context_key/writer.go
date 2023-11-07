package contextkey

import (
	"context"
)

type contextKey string

const (
	// ContextKeyUserID ユーザーを識別するID。認証後にコンテキストにセットされる
	ContextKeyUserID contextKey = "ctx-user-id"
)

type IContextWriter interface {
	SetUserID(ctx context.Context, userID string) context.Context
}

type ContextWriter struct{}

func NewContextWriter() IContextWriter {
	return &ContextWriter{}
}

func (w *ContextWriter) SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}
