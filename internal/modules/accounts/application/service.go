package application

import (
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/accounts/domain/ports"
)

type AccountService struct {
	accounts ports.AccountRepository
}

func NewAccountService(accounts ports.AccountRepository) *AccountService {
	return &AccountService{accounts: accounts}
}

func (s *AccountService) Create(name string, cpf string, agency string, deposit int64) error {

	err := s.accounts.Save(name, cpf, agency, deposit)
	if err != nil {
		return err
	}

	return nil
}
