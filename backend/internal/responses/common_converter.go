package responses

import (
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func NewIdModelFromJSONToProto(model *Id) *pb.Id {
	return &pb.Id{
		Id:        model.Id,
		AuthToken: model.AuthToken,
	}
}
func NewIdModelFromProtoToJSON(model *pb.Id) *Id {
	return &Id{
		Id:        model.GetId(),
		AuthToken: model.GetAuthToken(),
	}
}

func NewStatusResponseModelFromJSONToProto(model *StatusResponse) *pb.StatusResponse {
	return &pb.StatusResponse{
		Message: model.Message,
	}
}
func NewStatusResponseModelFromProtoToJSON(model *pb.StatusResponse) *StatusResponse {
	return &StatusResponse{
		Message: model.GetMessage(),
	}
}
