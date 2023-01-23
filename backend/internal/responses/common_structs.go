package responses

type StatusResponse struct {
	Message string `json:"message,omitempty"`
}

type Id struct {
	Id        int64  `json:"id,omitempty"`
	AuthToken string `json:"auth_token,omitempty"`
}
