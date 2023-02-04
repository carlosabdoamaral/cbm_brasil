package grpc

import (
	"context"

	persistence "github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence/occurrence"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *OccurrenceServer) Create(ctx context.Context, req *pb.CreateOccurrence) (*pb.StatusResponse, error) {
	err := persistence.CreateOccurrenceHandler(ctx, req)
	if err != nil {
		return &pb.StatusResponse{
			Message: err.Error(),
		}, err
	}

	return &pb.StatusResponse{
		Message: "created successfully",
	}, nil
}

func (s *OccurrenceServer) GetById(ctx context.Context, req *pb.Id) (*pb.OccurrenceDetails, error) {
	return persistence.GetOccurreceByIdHandler(ctx, req)
}

func (s *OccurrenceServer) GetAll(ctx context.Context, req *pb.Id) (*pb.OccurrenceDetailsList, error) {
	return persistence.GetAllOccurrences(ctx)
}

func (s *OccurrenceServer) AcceptById(ctx context.Context, req *pb.UpdateOccurrenceStatus) (*pb.StatusResponse, error) {
	return persistence.AcceptOccurenceByIdHandler(ctx, req)
}

func (s *OccurrenceServer) RefuseById(ctx context.Context, req *pb.UpdateOccurrenceStatus) (*pb.StatusResponse, error) {
	// return persistence.(ctx, req)
	//TODO: Implements this function
	return nil, nil
}

func (s *OccurrenceServer) CancelById(ctx context.Context, req *pb.UpdateOccurrenceStatus) (*pb.StatusResponse, error) {
	return persistence.CancelOccurenceByIdHandler(ctx, req)
}

func (s *OccurrenceServer) FinishById(ctx context.Context, req *pb.UpdateOccurrenceStatus) (*pb.StatusResponse, error) {
	return persistence.FinishOccurenceByIdHandler(ctx, req)
}
