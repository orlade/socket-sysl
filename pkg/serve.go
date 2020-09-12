package internal

import (
	"context"
	"fmt"
	"github.com/mlsquires/socketio"
	"log"
	"net/http"
)

// Serve starts listening for inbound requests to the service.
func Serve(
	ctx context.Context,
	cfg Config,
	newServer func(cfg Config, serviceIntf interface{}) (interface{}, error),
) error {
	ns, err := newServer(cfg, nil)
	if err != nil {
		return err
	}

	switch s := ns.(type) {
	case *socketio.Server:
		http.Handle("/socket.io/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			enableCors(w, r)
			s.ServeHTTP(w, r)
		}))

		listen := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		log.Printf("Serving on %s\n", listen)
		return http.ListenAndServe(listen, nil)

	default:
		panic(fmt.Errorf("unknown server type %T", s))
	}
}

// enableCors allows calls from any origin.
func enableCors(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	if r.Method == "OPTIONS" {
		return
	}
}
