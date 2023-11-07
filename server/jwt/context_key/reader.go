package contextkey

import (
	"context"
	"errors"
	interfaces "server/interfaces/context_key"
)

type ContextReader struct{}

func NewContextReader() interfaces.IContextReader {
	return &ContextReader{}
}

func (r *ContextReader) GetUserID(ctx context.Context) (string, error) {
	if v := ctx.Value(ContextKeyUserID); v != nil {
		if userID, ok := v.(string); ok {
			return userID, nil
		}
		return "", errors.New("error: context value not of type string for user-id")
	}
	return "", errors.New("error: context value not found for user-id")
}
