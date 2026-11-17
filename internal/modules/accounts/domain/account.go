package domain

type Account struct {
	ID      string
	Name    string
	CPF     string
	Agency  string
	Deposit int64
}

func (a *Account) IsValidName(name string) error {
	if name == "" {
		return ErrInvalidName
	}
	return nil
}

func (a *Account) IsValidCpf(cpf string) error {
	if cpf == "" {
		return ErrInvalidCPF
	}
	for _, r := range cpf {
		if r < '0' || r > '9' {
			return ErrInvalidCPF
		}
	}
	return nil
}

func (a *Account) IsValidAgency(agency string) error {
	if agency == "" {
		return ErrInvalidAgency
	}
	for _, r := range agency {
		if r < '0' || r > '9' {
			return ErrInvalidAgency
		}
	}
	return nil
}

func (a *Account) IsValidDeposit(deposit int64) error {
	if deposit <= 0 {
		return ErrInvalidInitialDeposit
	}
	return nil
}
