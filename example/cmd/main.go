package main

import (
	"context"
	exampleservice "github.com/orlade/socket-sysl/example/gen/ExampleService"
	"log"
)

func main() {
	//ch := func(sendMotd server.OnExampleServiceConnect) {
	//	log.Println("connect callback")
	//	sendMotd("hello")
	//}
	//
	//err := server.NewServer("localhost", 3001).
	//	HandleAll(map[string]interface{}{
	//		"connectz": ch,
	//	}).
	//	ListenAndServe()
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	log.Fatal(exampleservice.Serve(
		context.Background(),
		func(ctx context.Context, config interface{}) (*exampleservice.ServiceInterface, error) {
			return &exampleservice.ServiceInterface{
				OnExampleServiceConnect: func(ctx context.Context) error {
					log.Println("OnExampleServiceConnect")
					return nil
				},
			}, nil
		},
	))
}
