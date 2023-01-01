package models

import (
	"time"
)

type AccountPublicInfos struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type AccountPrivateInfos struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Location  Location  `json:"location,omitempty"`
	Token     Token     `json:"token,omitempty"`
	Password  string    `json:"password,omitempty"`
	IsDeleted bool      `json:"is_deleted,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Token struct {
	Code      string    `json:"code,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Location struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
