package exampleservice

import (
    "context"
)

// ServiceInterface for ExampleService.
type ServiceInterface struct {
    OnExampleServiceConnect func(ctx context.Context) error
}