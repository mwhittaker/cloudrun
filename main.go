package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

type server struct {
	weaver.Implements[weaver.Main]
	lis weaver.Listener
}

func serve(ctx context.Context, s *server) error {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	s.Logger(ctx).Info(fmt.Sprintf("Server listening on %v...", s.lis))
	return http.Serve(s.lis, nil)
}

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		panic(err)
	}
}
