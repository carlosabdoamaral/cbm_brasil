package responses

import (
	"testing"

	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
	"github.com/stretchr/testify/assert"
)

func TestNewSignInByEmailModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &SignInByEmail{
		Email:    "TEST@TEST.com",
		Password: "TEST",
	}

	protoModel := NewSignInByEmailModelFromJSONToProto(jsonModel)
	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.Email, protoModel.GetEmail())
	assert.Equal(t, jsonModel.Password, protoModel.GetPassword())
}
func TestNewSignInByEmailModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.SignInByEmailRequest{
		Email:    "TEST@TEST.com",
		Password: "TEST",
	}

	jsonModel := NewSignInByEmailModelFromProtoToJSON(protoModel)
	assert.NotEmpty(t, jsonModel)
	assert.Equal(t, protoModel.GetEmail(), jsonModel.Email)
	assert.Equal(t, protoModel.GetPassword(), jsonModel.Password)
}

func TestNewSignInByCPFModelFromJSONToProtoSuccess(t *testing.T) {
	jsonModel := &SignInByCPF{
		CPF:      "TEST@TEST.com",
		Password: "TEST",
	}

	protoModel := NewSignInByCPFModelFromJSONToProto(jsonModel)
	assert.NotEmpty(t, protoModel)
	assert.Equal(t, jsonModel.CPF, protoModel.GetCpf())
	assert.Equal(t, jsonModel.Password, protoModel.GetPassword())
}

func TestNewSignInByCPFModelFromProtoToJSONSuccess(t *testing.T) {
	protoModel := &pb.SignInByCPFRequest{
		Cpf:      "TEST@TEST.com",
		Password: "TEST",
	}

	jsonModel := NewSignInByCPFModelFromProtoToJSON(protoModel)
	assert.NotEmpty(t, jsonModel)
	assert.Equal(t, protoModel.GetCpf(), jsonModel.CPF)
	assert.Equal(t, protoModel.GetPassword(), jsonModel.Password)
}
