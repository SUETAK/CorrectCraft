package interfaces

import "server/domain/open_ai"

type OpenAI interface {
	GetCompletion(prompt string, maxTokens int) (open_ai.GptResponse, error)
}
