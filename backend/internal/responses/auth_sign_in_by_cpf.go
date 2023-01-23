package responses

import pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"

type SignInByCPFRequest struct {
	Cpf      string `json:"cpf,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewSignInByCPFRequestFromProtoToJSON(protoMessage *pb.SignInByCPFRequest) *SignInByCPFRequest {
	return &SignInByCPFRequest{
		Cpf:      protoMessage.GetCpf(),
		Password: protoMessage.GetPassword(),
	}
}

func NewSignInByCPFRequestFromJSONToProto(jsonMessage *SignInByCPFRequest) *pb.SignInByCPFRequest {
	return &pb.SignInByCPFRequest{
		Cpf:      jsonMessage.Cpf,
		Password: jsonMessage.Password,
	}
}
