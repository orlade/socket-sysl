package exampleservice

import (
    "context"
    "fmt"
    "github.com/mlsquires/socketio"
    internal "github.com/orlade/socket-sysl/pkg"
    "log"
)

// Serve starts the server.
//
// appConfig is a type defined by the application programmer to
// hold application-level configuration.
func Serve(
    ctx context.Context,
    createService func(ctx context.Context, appConfig interface{}) (*ServiceInterface, error),
) error {
    return internal.Serve(
        ctx,
        internal.Config{Host: "localhost", Port: 3001},
        func(cfg internal.Config, serviceIntf interface{}) (interface{}, error) {
            ss, err := socketio.NewServer(nil)
            if err != nil {
                return nil, err
            }

            svc, err := createService(ctx, cfg)
            if err != nil {
                return nil, err
            }

            ss.On("connection", func(so socketio.Socket) {
                log.Println("user connected")

                _ = so.On("Connect", func(msg map[string]interface{}) {
    _ = svc.OnExampleServiceConnect(ctx)
})
            })

            ss.On("error", func(so socketio.Socket, err error) {
                fmt.Printf("Error: %s\n", err)
            })

            return ss, nil
        },
    )
}