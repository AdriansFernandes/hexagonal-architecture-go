package httpapi

import "net/http"

func (h *Handle) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /account", h.create)
}
