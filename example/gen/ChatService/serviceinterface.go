package chatservice

import (
	"context"
)

// ServiceInterface for ChatService.
// Implementations of these methods must be provided.
type ServiceInterface struct {
	OnChatClientSendMessage func(ctx context.Context, msg map[string]interface{}, relayChatServiceSendMessage func(args ...interface{})) error
	OnChatServiceConnect    func(ctx context.Context, msg map[string]interface{}, emitMOTD func(args ...interface{})) error
}
