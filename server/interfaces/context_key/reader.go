package interfaces

import "context"

type IContextReader interface {
	GetUserID(ctx context.Context) (string, error)
}
