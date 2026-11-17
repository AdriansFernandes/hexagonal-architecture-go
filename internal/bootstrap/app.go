package bootstrap

import (
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/adapters/persistence/memory"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/application"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/entrypoints/httpapi"

	accountsmemory "github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/adapters/persistence/memory"
	accountsapp "github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/application"
	accountshttp "github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/entrypoints/httpapi"
)

func Bootstrap() http.Handler {
	accounts := memory.NewAccounts(
		domain.Account{ID: "a", Balance: 5000},
		domain.Account{ID: "b", Balance: 1000},
	)
	transfer := application.NewTransferService(accounts)
	handler := httpapi.NewHandler(transfer)

	accountRepository := accountsmemory.NewAccounts()
	accountService := accountsapp.NewAccountService(accountRepository)
	accountsHandler := accountshttp.NewHandler(accountService)

	routeHandlers := []RouteRegister{handler, accountsHandler}
	return NewRouter(routeHandlers...)
}
