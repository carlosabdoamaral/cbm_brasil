package grpc

import (
	"context"

	persistence "github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence/occurrence"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *OccurrenceServer) Create(ctx context.Context, req *pb.CreateOccurrence) (*pb.StatusResponse, error) {
	err := persistence.CreateOccurrenceHandler(&ctx, req)
	if err != nil {
		return &pb.StatusResponse{
			Message: err.Error(),
		}, err
	}

	return &pb.StatusResponse{
		Message: "created successfully",
	}, nil
}
