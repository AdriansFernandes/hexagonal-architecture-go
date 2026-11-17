package httpapi

import (
	"errors"
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/domain"
)

func statusFor(err error) int {
	switch {
	case errors.Is(err, domain.ErrInvalidName):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidCPF):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidAgency):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidInitialDeposit):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
