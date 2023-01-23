package responses

import pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"

type SignInByEmailRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewSignInByEmailRequestFromProtoToJSON(protoMessage *pb.SignInByEmailRequest) *SignInByEmailRequest {
	return &SignInByEmailRequest{
		Email:    protoMessage.GetEmail(),
		Password: protoMessage.GetPassword(),
	}
}

func NewSignInByEmailRequestFromJSONToProto(jsonMessage *SignInByEmailRequest) *pb.SignInByEmailRequest {
	return &pb.SignInByEmailRequest{
		Email:    jsonMessage.Email,
		Password: jsonMessage.Password,
	}
}
