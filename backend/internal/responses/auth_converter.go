package responses

import (
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

// Sign In By Email
func NewSignInByEmailModelFromJSONToProto(model *SignInByEmail) *pb.SignInByEmailRequest {
	return &pb.SignInByEmailRequest{
		Email:    model.Email,
		Password: model.Password,
	}
}
func NewSignInByEmailModelFromProtoToJSON(model *pb.SignInByEmailRequest) *SignInByEmail {
	return &SignInByEmail{
		Email:    model.GetEmail(),
		Password: model.GetPassword(),
	}
}

// Sign In By CPF
func NewSignInByCPFModelFromJSONToProto(model *SignInByCPF) *pb.SignInByCPFRequest {
	return &pb.SignInByCPFRequest{
		Cpf:      model.CPF,
		Password: model.Password,
	}
}
func NewSignInByCPFModelFromProtoToJSON(model *pb.SignInByCPFRequest) *SignInByCPF {
	return &SignInByCPF{
		CPF:      model.GetCpf(),
		Password: model.GetPassword(),
	}
}
