package responses

import (
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

type AccountSoftDeleteRequestJSON struct {
	Id        int64  `json:"id,omitempty"`
	AuthToken string `json:"auth_token,omitempty"`
}

func NewAccountSoftDeleteRequestFromProtoToJSON(protoMessage *pb.AccountSoftDeleteByIdRequest) *AccountSoftDeleteRequestJSON {
	return &AccountSoftDeleteRequestJSON{
		Id:        protoMessage.GetId(),
		AuthToken: protoMessage.GetAuthToken(),
	}
}

func NewAccountSoftDeleteRequestFromJSONToProto(jsonMessage *AccountSoftDeleteRequestJSON) *pb.AccountSoftDeleteByIdRequest {
	return &pb.AccountSoftDeleteByIdRequest{
		Id:        jsonMessage.Id,
		AuthToken: jsonMessage.AuthToken,
	}
}
