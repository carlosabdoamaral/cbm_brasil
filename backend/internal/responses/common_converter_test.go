package responses

import (
	"testing"

	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"github.com/stretchr/testify/assert"
)

func TestNewIdModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &Id{
		Id:        1,
		AuthToken: "TEST",
	}

	protoModel := NewIdModelFromJSONToProto(jsonModel)
	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.Id, protoModel.GetId())
	assert.Equal(t, jsonModel.AuthToken, protoModel.GetAuthToken())
}

func TestNewIdModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.Id{
		Id:        1,
		AuthToken: "TEST",
	}

	jsonModel := NewIdModelFromProtoToJSON(protoModel)
	assert.NotEmpty(t, jsonModel)
	assert.Equal(t, protoModel.GetId(), jsonModel.Id)
	assert.Equal(t, protoModel.GetAuthToken(), jsonModel.AuthToken)
}

func TestNewStatusResponseModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &StatusResponse{
		Message: "TEST",
	}

	protoModel := NewStatusResponseModelFromJSONToProto(jsonModel)
	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.Message, protoModel.GetMessage())
}

func TestNewStatusResponseModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.StatusResponse{
		Message: "TEST",
	}

	jsonModel := NewStatusResponseModelFromProtoToJSON(protoModel)
	assert.NotEmpty(t, jsonModel)
	assert.Equal(t, protoModel.GetMessage(), jsonModel.Message)
}
