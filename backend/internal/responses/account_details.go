package responses

import (
	"time"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

type AccountDetailsJSON struct {
	Id            int64               `json:"id,omitempty"`
	FullName      string              `json:"full_name,omitempty"`
	Email         string              `json:"email,omitempty"`
	Cpf           string              `json:"cpf,omitempty"`
	BirthDate     time.Time           `json:"birth_date,omitempty"`
	Password      string              `json:"password,omitempty"`
	TwoFactorCode string              `json:"two_factor_code,omitempty"`
	Location      LocationDetailsJSON `json:"location,omitempty"`
	CreatedAt     time.Time           `json:"created_at,omitempty"`
	UpdatedAt     time.Time           `json:"updated_at,omitempty"`
	SoftDeleted   bool                `json:"soft_deleted,omitempty"`
}

type LocationDetailsJSON struct {
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

func NewAccountDetailsFromProtoToJSON(protoMessage *pb.AccountDetails) *AccountDetailsJSON {
	return &AccountDetailsJSON{
		Id:            protoMessage.GetId(),
		FullName:      protoMessage.GetFullName(),
		Email:         protoMessage.GetEmail(),
		Cpf:           protoMessage.GetCpf(),
		BirthDate:     common.TimestampToTime(protoMessage.GetBirthDate()),
		Password:      protoMessage.GetPasswd(),
		TwoFactorCode: protoMessage.GetTwoFactorCode(),
		Location: LocationDetailsJSON{
			Id:           protoMessage.Location.GetId(),
			IdAccount:    protoMessage.Location.GetIdAccount(),
			CEP:          protoMessage.Location.GetCep(),
			Country:      protoMessage.Location.GetCountry(),
			State:        protoMessage.Location.GetState(),
			City:         protoMessage.Location.GetCity(),
			Neighborhood: protoMessage.Location.GetNeighborhood(),
			Street:       protoMessage.Location.GetStreet(),
			PlaceNumber:  protoMessage.Location.GetPlaceNumber(),
			Complement:   protoMessage.Location.GetComplement(),
		},
		CreatedAt:   common.TimestampToTime(protoMessage.GetCreatedAt()),
		UpdatedAt:   common.TimestampToTime(protoMessage.GetUpdatedAt()),
		SoftDeleted: protoMessage.GetSoftDeleted(),
	}
}

func NewAccountDetailsFromJSONToProto(jsonMessage *AccountDetailsJSON) *pb.AccountDetails {
	return &pb.AccountDetails{
		Id:            jsonMessage.Id,
		FullName:      jsonMessage.FullName,
		Email:         jsonMessage.Email,
		Cpf:           jsonMessage.Cpf,
		BirthDate:     common.TimeToTimestamp(jsonMessage.BirthDate),
		Passwd:        jsonMessage.Password,
		TwoFactorCode: jsonMessage.TwoFactorCode,
		Location: &pb.LocationDetails{
			Id:           jsonMessage.Location.Id,
			IdAccount:    jsonMessage.Location.IdAccount,
			Cep:          jsonMessage.Location.CEP,
			Country:      jsonMessage.Location.Country,
			State:        jsonMessage.Location.State,
			City:         jsonMessage.Location.City,
			Neighborhood: jsonMessage.Location.Neighborhood,
			Street:       jsonMessage.Location.Street,
			PlaceNumber:  jsonMessage.Location.PlaceNumber,
			Complement:   jsonMessage.Location.Complement,
		},
		CreatedAt:   common.TimeToTimestamp(jsonMessage.CreatedAt),
		UpdatedAt:   common.TimeToTimestamp(jsonMessage.UpdatedAt),
		SoftDeleted: jsonMessage.SoftDeleted,
	}
}
