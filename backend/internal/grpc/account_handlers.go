package grpc

import (
	"context"

	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *AccountServer) Create(c context.Context, req *pb.NewAccountRequest) (*pb.JwtToken, error) {
	return &pb.JwtToken{
		Token: "Success!",
	}, nil
}
