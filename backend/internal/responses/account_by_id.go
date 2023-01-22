package responses

import (
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

type AccountDetailsByIDRequestJSON struct {
	Id        int64  `json:"id,omitempty"`
	AuthToken string `json:"auth_token,omitempty"`
}

func NewAccountDetailsByIDRequestFromProtoToJSON(protoMessage *pb.GetAccountByIdRequest) *AccountDetailsByIDRequestJSON {
	return &AccountDetailsByIDRequestJSON{
		Id:        protoMessage.GetId(),
		AuthToken: protoMessage.GetAuthToken(),
	}
}

func NewAccountDetailsByIDRequestFromJSONToProto(jsonMessage *AccountDetailsByIDRequestJSON) *pb.GetAccountByIdRequest {
	return &pb.GetAccountByIdRequest{
		Id:        jsonMessage.Id,
		AuthToken: jsonMessage.AuthToken,
	}
}
