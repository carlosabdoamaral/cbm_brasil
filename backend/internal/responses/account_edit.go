package responses

import (
	"time"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

type EditAccountJSON struct {
	Id        int64                   `json:"id,omitempty"`
	FullName  string                  `json:"full_name,omitempty"`
	Email     string                  `json:"email,omitempty"`
	Cpf       string                  `json:"cpf,omitempty"`
	BirthDate time.Time               `json:"birth_date,omitempty"`
	Password  string                  `json:"password,omitempty"`
	Location  EditAccountLocationJSON `json:"location,omitempty"`
}

type EditAccountLocationJSON struct {
	CEP          string `json:"cep,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
	PlaceNumber  int64  `json:"place_number,omitempty"`
	Complement   string `json:"complement,omitempty"`
}

func NewEditAccountBodyFromProtoToJSON(protoMessage *pb.EditAccountByIdRequest) *EditAccountJSON {
	return &EditAccountJSON{
		Id:        protoMessage.GetId(),
		FullName:  protoMessage.GetFullName(),
		Email:     protoMessage.GetEmail(),
		Cpf:       protoMessage.GetCpf(),
		BirthDate: common.TimestampToTime(protoMessage.GetBirthDate()),
		Password:  protoMessage.GetPasswd(),
		Location: EditAccountLocationJSON{
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

func NewEditAccountBodyFromJSONToProto(jsonMessage *EditAccountJSON) *pb.EditAccountByIdRequest {
	return &pb.EditAccountByIdRequest{
		Id:        jsonMessage.Id,
		FullName:  jsonMessage.FullName,
		Email:     jsonMessage.Email,
		Cpf:       jsonMessage.Cpf,
		BirthDate: common.TimeToTimestamp(jsonMessage.BirthDate),
		Passwd:    jsonMessage.Password,
		Location: &pb.EditAccountByIdLocation{
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
