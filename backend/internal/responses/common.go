package responses

import (
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

type JwtToken struct {
	Jwt string `json:"jwt,omitempty"`
}

func NewJwtTokenFromProtoToJSON(protoMessage *pb.JwtToken) *JwtToken {
	return &JwtToken{
		Jwt: protoMessage.GetToken(),
	}
}

func NewJwtTokenFromJSONToProto(jsonMessage *JwtToken) *pb.JwtToken {
	return &pb.JwtToken{
		Token: jsonMessage.Jwt,
	}
}
