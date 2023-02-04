package responses

import "time"

type CreateOccurrence struct {
	IdFirefighter int64                    `json:"id_firefighter,omitempty"`
	IdAuthor      int64                    `json:"id_author,omitempty"`
	Title         string                   `json:"title,omitempty"`
	Description   string                   `json:"description,omitempty"`
	CreatedAt     time.Time                `json:"created_at,omitempty"`
	BannerX64     string                   `json:"banner_x64,omitempty"`
	Location      CreateOccurrenceLocation `json:"location,omitempty"`
}
type CreateOccurrenceLocation struct {
	CEP          string  `json:"cep,omitempty"`
	Country      string  `json:"country,omitempty"`
	State        string  `json:"state,omitempty"`
	City         string  `json:"city,omitempty"`
	Neighborhood string  `json:"neighborhood,omitempty"`
	Street       string  `json:"street,omitempty"`
	PlaceNumber  int64   `json:"place_number,omitempty"`
	Complement   string  `json:"complement,omitempty"`
	Latitude     float32 `json:"latitude,omitempty"`
	Longitude    float32 `json:"longitude,omitempty"`
}

type OccurrenceDetails struct {
	IdOccurrence        int64                     `json:"id_occurrence,omitempty"`
	Title               string                    `json:"title,omitempty"`
	Description         string                    `json:"description,omitempty"`
	BannerX64           string                    `json:"banner_x64,omitempty"`
	SoftDeleted         bool                      `json:"soft_deleted,omitempty"`
	IdFirefighter       int64                     `json:"id_firefighter,omitempty"`
	FirefighterFullname string                    `json:"firefighter_fullname,omitempty"`
	FirefighterEmail    string                    `json:"firefighter_email,omitempty"`
	IdAuthor            int64                     `json:"id_author,omitempty"`
	AuthorFullName      string                    `json:"author_fullname,omitempty"`
	AuthorEmail         string                    `json:"author_email,omitempty"`
	CreatedAt           time.Time                 `json:"created_at,omitempty"`
	UpdatedAt           time.Time                 `json:"updated_at,omitempty"`
	AcceptedAt          time.Time                 `json:"accepted_at,omitempty"`
	CanceledAt          time.Time                 `json:"canceled_at,omitempty"`
	FinishedAt          time.Time                 `json:"finished_at,omitempty"`
	IsAccepted          bool                      `json:"is_accepted,omitempty"`
	IsCanceled          bool                      `json:"is_canceled,omitempty"`
	IsFinished          bool                      `json:"is_finished,omitempty"`
	Location            OccurrenceDetailsLocation `json:"location,omitempty"`
}
type OccurrenceDetailsLocation struct {
	Id           int64   `json:"id,omitempty"`
	CEP          string  `json:"cep,omitempty"`
	Country      string  `json:"country,omitempty"`
	State        string  `json:"state,omitempty"`
	City         string  `json:"city,omitempty"`
	Neighborhood string  `json:"neighborhood,omitempty"`
	Street       string  `json:"street,omitempty"`
	PlaceNumber  int64   `json:"place_number,omitempty"`
	Complement   string  `json:"complement,omitempty"`
	Latitude     float32 `json:"latitude,omitempty"`
	Longitude    float32 `json:"longitude,omitempty"`
}

type EditOccurrence struct {
	Id          int64                  `json:"id,omitempty"`
	IdAuthor    int64                  `json:"id_author,omitempty"`
	BannerX64   string                 `json:"banner_x64,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Location    EditOccurrenceLocation `json:"location,omitempty"`
}
type EditOccurrenceLocation struct {
	Id           int64   `json:"id,omitempty"`
	CEP          string  `json:"cep,omitempty"`
	Country      string  `json:"country,omitempty"`
	State        string  `json:"state,omitempty"`
	City         string  `json:"city,omitempty"`
	Neighborhood string  `json:"neighborhood,omitempty"`
	Street       string  `json:"street,omitempty"`
	PlaceNumber  int64   `json:"place_number,omitempty"`
	Complement   string  `json:"complement,omitempty"`
	Latitude     float32 `json:"latitude,omitempty"`
	Longitude    float32 `json:"longitude,omitempty"`
}

type UpdateOccurrenceStatus struct {
	IdOccurrence  int64 `json:"id_occurrence,omitempty"`
	IdFirefighter int64 `json:"id_firefighter,omitempty"`
}
