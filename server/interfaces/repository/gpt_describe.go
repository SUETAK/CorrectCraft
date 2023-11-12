package interfaces

import "context"

type GPTDescribeRepository interface {
	CreateGPTDescribe(ctx context.Context, answerId int64, describe string) error
}
