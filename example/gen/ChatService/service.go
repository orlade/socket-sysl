package chatservice

import (
	"context"
)

// Service is the interface for ChatService clients.
type Service interface {
	OnChatClientSendMessage(ctx context.Context) error
	OnChatServiceConnect(ctx context.Context) error
}
