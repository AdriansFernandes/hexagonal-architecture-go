package domain

type Account struct {
	ID      string
	Balance int64
}

func (a *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Deposit(amount int64) {
	a.Balance += amount
}
