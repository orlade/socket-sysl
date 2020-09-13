package main

import (
	"context"
	"log"

	exampleservice "github.com/orlade/socket-sysl/example/gen/ExampleService"
)

func main() {
	log.Fatal(exampleservice.Serve(
		context.Background(),
		func(ctx context.Context, config interface{}) (*exampleservice.ServiceInterface, error) {
			return &exampleservice.ServiceInterface{
				OnExampleServiceConnect: func(ctx context.Context, emitMOTD func(args ...interface{})) error {
					log.Println("OnExampleServiceConnect")
					emitMOTD("hello!")
					return nil
				},
			}, nil
		},
	))
}
