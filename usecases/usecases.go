package usecases

import "context"

type PlatformUC interface {
	// starting for chat
	Chat(ctx context.Context)
}
