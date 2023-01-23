package responses

type SignInByEmail struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignInByCPF struct {
	CPF      string `json:"cpf,omitempty"`
	Password string `json:"password,omitempty"`
}
