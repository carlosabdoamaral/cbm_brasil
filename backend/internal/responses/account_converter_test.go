package responses

import (
	"testing"
	"time"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewCreateAccountModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &CreateAccount{
		FullName:      "TEST",
		Email:         "TEST",
		Cpf:           "TEST",
		BirthDate:     time.Now(),
		Password:      "TEST",
		TwoFactorCode: "TEST",
		Location: CreateAccountLocation{
			CEP:          "TEST",
			Country:      "TEST",
			State:        "TEST",
			City:         "TEST",
			Neighborhood: "TEST",
			Street:       "TEST",
			PlaceNumber:  100,
			Complement:   "TEST",
		},
	}

	protoModel := NewCreateAccountModelFromJSONToProto(jsonModel)

	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.FullName, protoModel.GetFullName())
	assert.Equal(t, jsonModel.Email, protoModel.GetEmail())
	assert.Equal(t, jsonModel.Cpf, protoModel.GetCpf())
	assert.Equal(t, common.TimeToTimestamp(jsonModel.BirthDate), protoModel.GetBirthDate())
	assert.Equal(t, jsonModel.Password, protoModel.GetPasswd())
	assert.Equal(t, jsonModel.TwoFactorCode, protoModel.GetTwoFactorCode())
	assert.Equal(t, jsonModel.Location.CEP, protoModel.Location.GetCep())
	assert.Equal(t, jsonModel.Location.Country, protoModel.Location.GetCountry())
	assert.Equal(t, jsonModel.Location.State, protoModel.Location.GetState())
	assert.Equal(t, jsonModel.Location.City, protoModel.Location.GetCity())
	assert.Equal(t, jsonModel.Location.Neighborhood, protoModel.Location.GetNeighborhood())
	assert.Equal(t, jsonModel.Location.Street, protoModel.Location.GetStreet())
	assert.Equal(t, jsonModel.Location.PlaceNumber, protoModel.Location.GetPlaceNumber())
	assert.Equal(t, jsonModel.Location.Complement, protoModel.Location.GetComplement())
}

func TestNewCreateAccountModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.CreateAccount{
		FullName:      "TEST",
		Email:         "TEST",
		Cpf:           "TEST",
		BirthDate:     timestamppb.Now(),
		Passwd:        "TEST",
		TwoFactorCode: "TEST",
		Location: &pb.CreateAccountLocation{
			Cep:          "TEST",
			Country:      "TEST",
			State:        "TEST",
			City:         "TEST",
			Neighborhood: "TEST",
			Street:       "TEST",
			PlaceNumber:  100,
			Complement:   "TEST",
		},
	}

	jsonModel := NewCreateAccountModelFromProtoToJSON(protoModel)

	assert.NotEmpty(t, jsonModel)
	assert.Equal(t, protoModel.GetFullName(), jsonModel.FullName)
	assert.Equal(t, protoModel.GetEmail(), jsonModel.Email)
	assert.Equal(t, protoModel.GetCpf(), jsonModel.Cpf)
	assert.Equal(t, protoModel.GetBirthDate(), common.TimeToTimestamp(jsonModel.BirthDate))
	assert.Equal(t, protoModel.GetPasswd(), jsonModel.Password)
	assert.Equal(t, protoModel.GetTwoFactorCode(), jsonModel.TwoFactorCode)
	assert.Equal(t, protoModel.Location.GetCep(), jsonModel.Location.CEP)
	assert.Equal(t, protoModel.Location.GetCountry(), jsonModel.Location.Country)
	assert.Equal(t, protoModel.Location.GetState(), jsonModel.Location.State)
	assert.Equal(t, protoModel.Location.GetCity(), jsonModel.Location.City)
	assert.Equal(t, protoModel.Location.GetNeighborhood(), jsonModel.Location.Neighborhood)
	assert.Equal(t, protoModel.Location.GetStreet(), jsonModel.Location.Street)
	assert.Equal(t, protoModel.Location.GetPlaceNumber(), jsonModel.Location.PlaceNumber)
	assert.Equal(t, protoModel.Location.GetComplement(), jsonModel.Location.Complement)
}

func TestNewAccountDetailsModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &AccountDetails{
		Id:            1,
		FullName:      "TEST",
		Email:         "TEST",
		Cpf:           "TEST",
		BirthDate:     time.Now(),
		Password:      "TEST",
		TwoFactorCode: "TEST",
		SoftDeleted:   false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Location: AccountLocationDetails{
			Id:           1,
			IdAccount:    1,
			CEP:          "TEST",
			Country:      "TEST",
			State:        "TEST",
			City:         "TEST",
			Neighborhood: "TEST",
			Street:       "TEST",
			PlaceNumber:  1,
			Complement:   "TEST",
		},
	}

	protoModel := NewAccountDetailsModelFromJSONToProto(jsonModel)
	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.Id, protoModel.GetId())
	assert.Equal(t, jsonModel.FullName, protoModel.GetFullName())
	assert.Equal(t, jsonModel.Email, protoModel.GetEmail())
	assert.Equal(t, jsonModel.Cpf, protoModel.GetCpf())
	assert.Equal(t, jsonModel.BirthDate.UTC(), common.TimestampToTime(protoModel.GetBirthDate()).UTC())
	assert.Equal(t, jsonModel.Password, protoModel.GetPasswd())
	assert.Equal(t, jsonModel.TwoFactorCode, protoModel.GetTwoFactorCode())
	assert.Equal(t, jsonModel.SoftDeleted, protoModel.GetSoftDeleted())
	assert.Equal(t, jsonModel.CreatedAt.UTC(), common.TimestampToTime(protoModel.GetCreatedAt()).UTC())
	assert.Equal(t, jsonModel.UpdatedAt.UTC(), common.TimestampToTime(protoModel.GetUpdatedAt()).UTC())
	assert.Equal(t, jsonModel.Location.Id, protoModel.Location.GetId())
	assert.Equal(t, jsonModel.Location.IdAccount, protoModel.Location.GetIdAccount())
	assert.Equal(t, jsonModel.Location.CEP, protoModel.Location.GetCep())
	assert.Equal(t, jsonModel.Location.Country, protoModel.Location.GetCountry())
	assert.Equal(t, jsonModel.Location.State, protoModel.Location.GetState())
	assert.Equal(t, jsonModel.Location.City, protoModel.Location.GetCity())
	assert.Equal(t, jsonModel.Location.Neighborhood, protoModel.Location.GetNeighborhood())
	assert.Equal(t, jsonModel.Location.Street, protoModel.Location.GetStreet())
	assert.Equal(t, jsonModel.Location.PlaceNumber, protoModel.Location.GetPlaceNumber())
	assert.Equal(t, jsonModel.Location.Complement, protoModel.Location.GetComplement())
}

func TestNewAccountDetailsModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.AccountDetails{
		Id:            1,
		FullName:      "TEST",
		Email:         "TEST",
		Cpf:           "TEST",
		BirthDate:     timestamppb.Now(),
		Passwd:        "TEST",
		TwoFactorCode: "TEST",
		SoftDeleted:   false,
		CreatedAt:     timestamppb.Now(),
		UpdatedAt:     timestamppb.Now(),
		Location: &pb.AccountLocationDetails{
			Id:           1,
			IdAccount:    1,
			Cep:          "TEST",
			Country:      "TEST",
			State:        "TEST",
			City:         "TEST",
			Neighborhood: "TEST",
			Street:       "TEST",
			PlaceNumber:  1,
			Complement:   "TEST",
		},
	}

	jsonModel := NewAccountDetailsModelFromProtoToJSON(protoModel)

	assert.NotEmpty(t, protoModel)
	assert.Equal(t, protoModel.GetId(), jsonModel.Id)
	assert.Equal(t, protoModel.GetFullName(), jsonModel.FullName)
	assert.Equal(t, protoModel.GetEmail(), jsonModel.Email)
	assert.Equal(t, protoModel.GetCpf(), jsonModel.Cpf)
	assert.Equal(t, protoModel.GetBirthDate().AsTime().UTC(), jsonModel.BirthDate.UTC())
	assert.Equal(t, protoModel.GetPasswd(), jsonModel.Password)
	assert.Equal(t, protoModel.GetTwoFactorCode(), jsonModel.TwoFactorCode)
	assert.Equal(t, protoModel.GetSoftDeleted(), jsonModel.SoftDeleted)
	assert.Equal(t, protoModel.GetCreatedAt().AsTime().UTC(), jsonModel.CreatedAt.UTC())
	assert.Equal(t, protoModel.GetUpdatedAt().AsTime().UTC(), jsonModel.UpdatedAt.UTC())
	assert.Equal(t, protoModel.Location.GetId(), jsonModel.Location.Id)
	assert.Equal(t, protoModel.Location.GetIdAccount(), jsonModel.Location.IdAccount)
	assert.Equal(t, protoModel.Location.GetCep(), jsonModel.Location.CEP)
	assert.Equal(t, protoModel.Location.GetCountry(), jsonModel.Location.Country)
	assert.Equal(t, protoModel.Location.GetState(), jsonModel.Location.State)
	assert.Equal(t, protoModel.Location.GetCity(), jsonModel.Location.City)
	assert.Equal(t, protoModel.Location.GetNeighborhood(), jsonModel.Location.Neighborhood)
	assert.Equal(t, protoModel.Location.GetStreet(), jsonModel.Location.Street)
	assert.Equal(t, protoModel.Location.GetPlaceNumber(), jsonModel.Location.PlaceNumber)
	assert.Equal(t, protoModel.Location.GetComplement(), jsonModel.Location.Complement)
}

func TestNewEditAccountModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &EditAccount{
		Id:        1,
		FullName:  "TEST",
		Email:     "TEST",
		Cpf:       "TEST",
		BirthDate: time.Now(),
		Password:  "TEST",
		Location: EditAccountLocation{
			CEP:          "TEST",
			Country:      "TEST",
			State:        "TEST",
			City:         "TEST",
			Neighborhood: "TEST",
			Street:       "TEST",
			PlaceNumber:  1,
			Complement:   "TEST",
		},
	}

	protoModel := NewEditAccountModelFromJSONToProto(jsonModel)

	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.Id, protoModel.GetId())
	assert.Equal(t, jsonModel.FullName, protoModel.GetFullName())
	assert.Equal(t, jsonModel.Email, protoModel.GetEmail())
	assert.Equal(t, jsonModel.Cpf, protoModel.GetCpf())
	assert.Equal(t, jsonModel.BirthDate.UTC(), common.TimestampToTime(protoModel.GetBirthDate()).UTC())
	assert.Equal(t, jsonModel.Password, protoModel.GetPasswd())
	assert.Equal(t, jsonModel.Location.CEP, protoModel.Location.GetCep())
	assert.Equal(t, jsonModel.Location.Country, protoModel.Location.GetCountry())
	assert.Equal(t, jsonModel.Location.State, protoModel.Location.GetState())
	assert.Equal(t, jsonModel.Location.City, protoModel.Location.GetCity())
	assert.Equal(t, jsonModel.Location.Neighborhood, protoModel.Location.GetNeighborhood())
	assert.Equal(t, jsonModel.Location.Street, protoModel.Location.GetStreet())
	assert.Equal(t, jsonModel.Location.PlaceNumber, protoModel.Location.GetPlaceNumber())
	assert.Equal(t, jsonModel.Location.Complement, protoModel.Location.GetComplement())
}

func TestNewEditAccountModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.EditAccount{
		Id:        1,
		FullName:  "TEST",
		Email:     "TEST",
		Cpf:       "TEST",
		BirthDate: timestamppb.Now(),
		Passwd:    "TEST",
		Location: &pb.EditAccountLocation{
			Cep:          "TEST",
			Country:      "TEST",
			State:        "TEST",
			City:         "TEST",
			Neighborhood: "TEST",
			Street:       "TEST",
			PlaceNumber:  1,
			Complement:   "TEST",
		},
	}

	jsonModel := NewEditAccountModelFromProtoToJSON(protoModel)

	assert.NotEmpty(t, jsonModel)
	assert.Equal(t, protoModel.GetId(), jsonModel.Id)
	assert.Equal(t, protoModel.GetFullName(), jsonModel.FullName)
	assert.Equal(t, protoModel.GetEmail(), jsonModel.Email)
	assert.Equal(t, protoModel.GetCpf(), jsonModel.Cpf)
	assert.Equal(t, common.TimestampToTime(protoModel.GetBirthDate()).UTC(), jsonModel.BirthDate.UTC())
	assert.Equal(t, protoModel.GetPasswd(), jsonModel.Password)
	assert.Equal(t, protoModel.Location.GetCep(), jsonModel.Location.CEP)
	assert.Equal(t, protoModel.Location.GetCountry(), jsonModel.Location.Country)
	assert.Equal(t, protoModel.Location.GetState(), jsonModel.Location.State)
	assert.Equal(t, protoModel.Location.GetCity(), jsonModel.Location.City)
	assert.Equal(t, protoModel.Location.GetNeighborhood(), jsonModel.Location.Neighborhood)
	assert.Equal(t, protoModel.Location.GetStreet(), jsonModel.Location.Street)
	assert.Equal(t, protoModel.Location.GetPlaceNumber(), jsonModel.Location.PlaceNumber)
	assert.Equal(t, protoModel.Location.GetComplement(), jsonModel.Location.Complement)
}
