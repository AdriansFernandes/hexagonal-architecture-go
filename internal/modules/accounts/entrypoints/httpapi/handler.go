package httpapi

import (
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/application"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/platform/httpserver"
)

type Handle struct {
	service *application.AccountService
}

func NewHandler(service *application.AccountService) *Handle {
	return &Handle{service: service}
}

func (h *Handle) create(w http.ResponseWriter, r *http.Request) {
	var input accountResquest
	if err := httpserver.ReadJSON(w, r, &input); err != nil {
		httpserver.Fail(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	err := h.service.Create(input.Name, input.CPF, input.Agency, input.Deposit)
	if err != nil {
		httpserver.Fail(w, statusFor(err), err.Error())
		return
	}
}
