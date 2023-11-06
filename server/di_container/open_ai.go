package di

import (
	"server/domain/open_ai"
	interfaces "server/interfaces/open_ai"
)

func NewOpenAIClient(key string) interfaces.OpenAI {
	return &open_ai.OpenAIClient{
		Key: key,
	}
}
