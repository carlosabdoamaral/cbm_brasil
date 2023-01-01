package common

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtoToJwt(req protoreflect.ProtoMessage) ([]byte, error) {
	return protojson.Marshal(req)
}
