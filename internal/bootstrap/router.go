package bootstrap

import (
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/platform/httpserver"
)

type RouteRegister interface {
	RegisterRoutes(mux *http.ServeMux)
}

func NewRouter(handlers ...RouteRegister) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", httpserver.Health)

	for _, handler := range handlers {
		handler.RegisterRoutes(mux)
	}

	return mux
}
