package ports

type AccountRepository interface {
	Save(name string, cpf string, agency string, deposit int64) error
	// List(accounts []domain.Account) error
}
