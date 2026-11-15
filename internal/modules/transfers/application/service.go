package application

import "github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain/ports"

type TransferService struct {
	accounts ports.AccountRepository
}

func NewTransferService(accounts ports.AccountRepository) *TransferService {
	return &TransferService{accounts: accounts}
}

func (s *TransferService) Execute(fromID string, toID string, amount int64) error {
	from, err := s.accounts.Get(fromID)

	if err != nil {
		return err
	}

	to, err := s.accounts.Get(toID)
	if err != nil {
		return err
	}

	if err := from.Withdraw(amount); err != nil {
		return err
	}

	to.Deposit(amount)

	return s.accounts.Save(from, to)
}
