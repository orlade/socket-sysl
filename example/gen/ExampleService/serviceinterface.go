package exampleservice

import (
	"context"
)

// ServiceInterface for ExampleService.
// Implementations of these methods must be provided.
type ServiceInterface struct {
	OnExampleServiceConnect func(ctx context.Context, emitMOTD func(args ...interface{})) error
}
