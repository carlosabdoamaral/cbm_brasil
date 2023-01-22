package responses

import (
	"time"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NewAccountRequestJSON struct {
	FullName      string                        `json:"full_name,omitempty"`
	Email         string                        `json:"email,omitempty"`
	Cpf           string                        `json:"cpf,omitempty"`
	BirthDate     time.Time                     `json:"birth_date,omitempty"`
	Password      string                        `json:"password,omitempty"`
	TwoFactorCode string                        `json:"two_factor_code,omitempty"`
	Location      NewAccountRequestLocationJSON `json:"location,omitempty"`
}

type NewAccountRequestLocationJSON struct {
	CEP          string `json:"cep,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
	PlaceNumber  int64  `json:"place_number,omitempty"`
	Complement   string `json:"complement,omitempty"`
}

func NewAccountRequestFromProtoToJSON(protoMessage *pb.NewAccountRequest) *NewAccountRequestJSON {
	return &NewAccountRequestJSON{
		FullName:      protoMessage.GetFullName(),
		Email:         protoMessage.GetEmail(),
		Cpf:           protoMessage.GetCpf(),
		BirthDate:     common.TimestampToTime(protoMessage.GetBirthDate()),
		Password:      protoMessage.GetPasswd(),
		TwoFactorCode: protoMessage.GetTwoFactorCode(),
		Location: NewAccountRequestLocationJSON{
			CEP:          protoMessage.Location.GetCep(),
			Country:      protoMessage.Location.GetCountry(),
			State:        protoMessage.Location.GetState(),
			City:         protoMessage.Location.GetCity(),
			Neighborhood: protoMessage.Location.GetNeighborhood(),
			Street:       protoMessage.Location.GetStreet(),
			PlaceNumber:  protoMessage.Location.GetPlaceNumber(),
			Complement:   protoMessage.Location.GetComplement(),
		},
	}
}

func NewAccountRequestFromJSONToProto(jsonMessage *NewAccountRequestJSON) *pb.NewAccountRequest {
	return &pb.NewAccountRequest{
		FullName:      jsonMessage.FullName,
		Email:         jsonMessage.Email,
		Cpf:           jsonMessage.Cpf,
		BirthDate:     timestamppb.Now(), // TODO: Adicionar -> common.TimeToTimestamp(jsonMessage.BirthDate)
		Passwd:        jsonMessage.Password,
		TwoFactorCode: jsonMessage.TwoFactorCode,
		Location: &pb.NewAccountLocation{
			Cep:          jsonMessage.Location.CEP,
			Country:      jsonMessage.Location.Country,
			State:        jsonMessage.Location.State,
			City:         jsonMessage.Location.City,
			Neighborhood: jsonMessage.Location.Neighborhood,
			Street:       jsonMessage.Location.Street,
			PlaceNumber:  jsonMessage.Location.PlaceNumber,
			Complement:   jsonMessage.Location.Complement,
		},
	}
}
