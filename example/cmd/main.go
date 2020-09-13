package main

import (
	"context"
	"log"
	"net/http"

	chatservice "github.com/orlade/socket-sysl/example/gen/ChatService"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.Handle("/gen/", http.StripPrefix("/gen", http.FileServer(http.Dir("gen/ChatClient/"))))

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
