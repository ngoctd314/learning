package trainings

import (
	"context"
	"net/http"

	"ddd/internal/common/server"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()
	_ = ctx

	server.RunHTTPServer(func(router chi.Router) http.Handler { return nil })
}
