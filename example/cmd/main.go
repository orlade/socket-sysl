package main

import (
	"context"
	"log"

	chatservice "github.com/orlade/socket-sysl/example/gen/ChatService"
)

func main() {
	log.Fatal(chatservice.Serve(
		context.Background(),
		func(ctx context.Context, config interface{}) (*chatservice.ServiceInterface, error) {
			return &chatservice.ServiceInterface{
				OnChatServiceConnect: func(ctx context.Context, msg map[string]interface{}, emitMOTD func(args ...interface{})) error {
					log.Println("OnExampleServiceConnect")
					emitMOTD("hello!")
					return nil
				},
				OnChatClientSendMessage: func(ctx context.Context, msg map[string]interface{}, relayChatServiceSendMessage func(args ...interface{})) error {
					log.Println("OnChatClientSendMessage")
					relayChatServiceSendMessage(msg)
					return nil
				},
			}, nil
		},
	))
}
