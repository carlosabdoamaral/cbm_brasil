package responses

import "time"

// Create Account
type CreateAccount struct {
	FullName      string                `json:"full_name,omitempty"`
	Email         string                `json:"email,omitempty"`
	Cpf           string                `json:"cpf,omitempty"`
	BirthDate     time.Time             `json:"birth_date,omitempty"`
	Password      string                `json:"password,omitempty"`
	TwoFactorCode string                `json:"two_factor_code,omitempty"`
	Location      CreateAccountLocation `json:"location,omitempty"`
}

type CreateAccountLocation struct {
	CEP          string `json:"cep,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
	PlaceNumber  int64  `json:"place_number,omitempty"`
	Complement   string `json:"complement,omitempty"`
}

// Account Details
type AccountDetails struct {
	Id            int64                  `json:"id,omitempty"`
	FullName      string                 `json:"full_name,omitempty"`
	Email         string                 `json:"email,omitempty"`
	Cpf           string                 `json:"cpf,omitempty"`
	BirthDate     time.Time              `json:"birth_date,omitempty"`
	Password      string                 `json:"password,omitempty"`
	TwoFactorCode string                 `json:"two_factor_code,omitempty"`
	SoftDeleted   bool                   `json:"soft_deleted,omitempty"`
	CreatedAt     time.Time              `json:"created_at,omitempty"`
	UpdatedAt     time.Time              `json:"updated_at,omitempty"`
	Location      AccountLocationDetails `json:"location,omitempty"`
}

type AccountLocationDetails struct {
	Id           int64  `json:"id,omitempty"`
	IdAccount    int64  `json:"id_account,omitempty"`
	CEP          string `json:"cep,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
	PlaceNumber  int64  `json:"place_number,omitempty"`
	Complement   string `json:"complement,omitempty"`
}

// Edit Account
type EditAccount struct {
	Id        int64               `json:"id,omitempty"`
	FullName  string              `json:"full_name,omitempty"`
	Email     string              `json:"email,omitempty"`
	Cpf       string              `json:"cpf,omitempty"`
	BirthDate time.Time           `json:"birth_date,omitempty"`
	Password  string              `json:"password,omitempty"`
	Location  EditAccountLocation `json:"location,omitempty"`
}

type EditAccountLocation struct {
	CEP          string `json:"cep,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
	PlaceNumber  int64  `json:"place_number,omitempty"`
	Complement   string `json:"complement,omitempty"`
}
