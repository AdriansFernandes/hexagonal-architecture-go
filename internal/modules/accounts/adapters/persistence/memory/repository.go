package memory

import (
	"sync"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/domain/ports"
)

type Account struct {
	mu    sync.RWMutex
	items map[string]domain.Account
}

func NewAccounts(initial ...domain.Account) *Account {
	items := make(map[string]domain.Account, len(initial))

	for _, account := range initial {
		items[account.ID] = account
	}

	return &Account{items: items}
}

// Save implements [ports.AccountRepository].
func (a *Account) Save(name string, cpf string, agency string, deposit int64) error {
	return nil
}

var _ ports.AccountRepository = (*Account)(nil)
