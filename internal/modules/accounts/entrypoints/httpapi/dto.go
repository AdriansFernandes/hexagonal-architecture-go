package httpapi

type accountResquest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Agency  string `json:"agency"`
	Deposit int64  `json:"deposit"`
}
