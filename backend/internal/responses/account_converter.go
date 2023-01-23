package responses

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

// Create Account
func NewCreateAccountModelFromJSONToProto(model *CreateAccount) *pb.CreateAccount {
	return &pb.CreateAccount{
		FullName:      model.FullName,
		Email:         model.Email,
		Cpf:           model.Cpf,
		BirthDate:     common.TimeToTimestamp(model.BirthDate),
		Passwd:        model.Password,
		TwoFactorCode: model.TwoFactorCode,
		Location: &pb.CreateAccountLocation{
			Cep:          model.Location.CEP,
			Country:      model.Location.Country,
			State:        model.Location.State,
			City:         model.Location.City,
			Neighborhood: model.Location.Neighborhood,
			Street:       model.Location.Street,
			PlaceNumber:  model.Location.PlaceNumber,
			Complement:   model.Location.Complement,
		},
	}
}
func NewCreateAccountModelFromProtoToJSON(model *pb.CreateAccount) *CreateAccount {
	return &CreateAccount{
		FullName:      model.GetFullName(),
		Email:         model.GetEmail(),
		Cpf:           model.GetCpf(),
		BirthDate:     common.TimestampToTime(model.GetBirthDate()),
		Password:      model.GetPasswd(),
		TwoFactorCode: model.GetTwoFactorCode(),
		Location: CreateAccountLocation{
			CEP:          model.Location.GetCep(),
			Country:      model.Location.GetCountry(),
			State:        model.Location.GetState(),
			City:         model.Location.GetCity(),
			Neighborhood: model.Location.GetNeighborhood(),
			Street:       model.Location.GetStreet(),
			PlaceNumber:  model.Location.GetPlaceNumber(),
			Complement:   model.Location.GetComplement(),
		},
	}
}

// Account Details
func NewAccountDetailsModelFromJSONToProto(model *AccountDetails) *pb.AccountDetails {
	return &pb.AccountDetails{
		Id:            model.Id,
		FullName:      model.FullName,
		Email:         model.Email,
		Cpf:           model.Cpf,
		BirthDate:     common.TimeToTimestamp(model.BirthDate),
		Passwd:        model.Password,
		TwoFactorCode: model.TwoFactorCode,
		Location: &pb.AccountLocationDetails{
			Id:           model.Location.Id,
			IdAccount:    model.Location.IdAccount,
			Cep:          model.Location.CEP,
			Country:      model.Location.Country,
			State:        model.Location.State,
			City:         model.Location.Street,
			Neighborhood: model.Location.Neighborhood,
			Street:       model.Location.Street,
			PlaceNumber:  model.Location.PlaceNumber,
			Complement:   model.Location.Complement,
		},
		CreatedAt:   common.TimeToTimestamp(model.CreatedAt),
		UpdatedAt:   common.TimeToTimestamp(model.UpdatedAt),
		SoftDeleted: model.SoftDeleted,
	}
}
func NewAccountDetailsModelFromProtoToJSON(model *pb.AccountDetails) *AccountDetails {
	return &AccountDetails{
		Id:            model.GetId(),
		FullName:      model.GetFullName(),
		Email:         model.GetEmail(),
		Cpf:           model.GetCpf(),
		BirthDate:     common.TimestampToTime(model.GetBirthDate()),
		Password:      model.GetPasswd(),
		TwoFactorCode: model.GetTwoFactorCode(),
		SoftDeleted:   model.GetSoftDeleted(),
		CreatedAt:     common.TimestampToTime(model.CreatedAt),
		UpdatedAt:     common.TimestampToTime(model.UpdatedAt),
		Location: AccountLocationDetails{
			Id:           model.Location.GetId(),
			IdAccount:    model.Location.GetIdAccount(),
			CEP:          model.Location.GetCep(),
			Country:      model.Location.GetCountry(),
			State:        model.Location.GetState(),
			City:         model.Location.GetCity(),
			Neighborhood: model.Location.GetNeighborhood(),
			Street:       model.Location.GetStreet(),
			PlaceNumber:  model.Location.GetPlaceNumber(),
			Complement:   model.Location.GetComplement(),
		},
	}
}

// Edit Account
func NewEditAccountModelFromJSONToProto(model *EditAccount) *pb.EditAccount {
	return &pb.EditAccount{
		Id:        model.Id,
		FullName:  model.FullName,
		Email:     model.Email,
		Cpf:       model.Cpf,
		BirthDate: common.TimeToTimestamp(model.BirthDate),
		Passwd:    model.Password,
		Location: &pb.EditAccountLocation{
			Cep:          model.Location.CEP,
			Country:      model.Location.Country,
			State:        model.Location.State,
			City:         model.Location.City,
			Neighborhood: model.Location.Neighborhood,
			Street:       model.Location.Street,
			PlaceNumber:  model.Location.PlaceNumber,
			Complement:   model.Location.Complement,
		},
	}
}
func NewEditAccountModelFromProtoToJSON(model *pb.EditAccount) *EditAccount {
	return &EditAccount{
		Id:        model.GetId(),
		FullName:  model.GetFullName(),
		Email:     model.GetEmail(),
		Cpf:       model.GetCpf(),
		BirthDate: common.TimestampToTime(model.GetBirthDate()),
		Password:  model.GetPasswd(),
		Location: EditAccountLocation{
			CEP:          model.Location.GetCep(),
			Country:      model.Location.GetCountry(),
			State:        model.Location.GetState(),
			City:         model.Location.GetCity(),
			Neighborhood: model.Location.GetNeighborhood(),
			Street:       model.Location.GetStreet(),
			PlaceNumber:  model.Location.GetPlaceNumber(),
			Complement:   model.Location.GetComplement(),
		},
	}
}
