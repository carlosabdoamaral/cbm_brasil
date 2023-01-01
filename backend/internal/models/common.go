package models

type JwtToken struct {
	Jwt []byte `json:"jwt,omitempty"`
}
