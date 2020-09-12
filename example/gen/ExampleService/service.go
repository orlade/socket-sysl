package exampleservice

import (
    "context"
)

// Service is the interface for ExampleService clients.
type Service interface {
    OnExampleServiceConnect(ctx context.Context) error
}