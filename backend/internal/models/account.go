package models

import "google.golang.org/protobuf/types/known/timestamppb"

type AccountPublicInfos struct {
	Id        int64                  `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	CreatedAt *timestamppb.Timestamp `json:"createdAt,omitempty"`
}

type AccountPrivateInfos struct {
	Id        int64                  `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Email     string                 `json:"email,omitempty"`
	Location  Location               `json:"location,omitempty"`
	Token     Token                  `json:"token,omitempty"`
	Password  string                 `json:"password,omitempty"`
	IsDeleted bool                   `json:"is_deleted,omitempty"`
	CreatedAt *timestamppb.Timestamp `json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at,omitempty"`
}

type Token struct {
	Code      string                 `json:"code,omitempty"`
	CreatedAt *timestamppb.Timestamp `json:"created_at,omitempty"`
}

type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}
