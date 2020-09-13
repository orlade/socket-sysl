# socket-sysl example

An example socket-sysl app.

## Build

```bash
make
```

## Implementation

Code generation will produce all the moving parts that you need to build the system, but they still require some simple orchestration.

In your `main` function, you'll need to call a `Serve` function to serve the API that the client will use. This requires the construction of a `ServiceInterface`, which in turn requires the implementation of functions to handle inbound requests.

These functions must implement a generated signature (based on the Sysl spec), and they will be passed helper functions to fire downstream events. 

Optionally, you may also want to serve the directory of generated client code on the `/` path, so that it can be accessed from the same server.

Here is an example with a `nil` implementation of the `ServiceInterface` for brevity. See the [implementation of `example`](example/) for a more complete demonstration.

```go
package main

import (
	"context"
	"log"
	"net/http"

	chatservice "github.com/orlade/socket-sysl/example/gen/ChatService"
)

func main() {
	// Serve the generated client code from the base path.
	http.Handle("/", http.FileServer(http.Dir("gen/ChatClient")))

	// Listen and serve the WebSockets API.
	log.Fatal(chatservice.Serve(
		context.Background(),
		func(ctx context.Context, config interface{}) (*chatservice.ServiceInterface, error) {
			return &chatservice.ServiceInterface{
				// Provide implementations here.
				OnChatServiceConnect: nil,
				OnChatClientSendMessage: nil,
			 }, nil
		},
	))
}
```
