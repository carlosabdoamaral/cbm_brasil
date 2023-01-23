package grpc

import (
	"context"

	persistence "github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence/account"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *AccountServer) Create(ctx context.Context, req *pb.CreateAccount) (*pb.AccountDetails, error) {
	res, err := persistence.CreateAccount(&ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) GetById(ctx context.Context, req *pb.Id) (*pb.AccountDetails, error) {
	res, err := persistence.GetAccountDetailsById(&ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) EditById(ctx context.Context, req *pb.EditAccount) (*pb.AccountDetails, error) {
	res, err := persistence.EditAccountById(&ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) SoftDeleteById(ctx context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	res, err := persistence.UpdateSoftDeleteStateById(&ctx, true, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) RecoverAccountById(ctx context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	res, err := persistence.UpdateSoftDeleteStateById(&ctx, false, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
