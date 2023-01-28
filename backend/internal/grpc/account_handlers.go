package grpc

import (
	"context"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	persistence "github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence/account"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *AccountServer) Create(ctx context.Context, req *pb.CreateAccount) (*pb.AccountDetails, error) {
	common.LogInfo("AccountServer - Create")
	res, err := persistence.CreateAccount(&ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) GetById(ctx context.Context, req *pb.Id) (*pb.AccountDetails, error) {
	common.LogInfo("AccountServer - GetById")
	res, err := persistence.GetAccountDetailsById(&ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) EditById(ctx context.Context, req *pb.EditAccount) (*pb.AccountDetails, error) {
	common.LogInfo("AccountServer - EditById")
	res, err := persistence.EditAccountById(&ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) SoftDeleteById(ctx context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	common.LogInfo("AccountServer - SoftDeleteById")

	res, err := persistence.UpdateSoftDeleteStateById(&ctx, true, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) RecoverAccountById(ctx context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	common.LogInfo("AccountServer - RecoverAccountById")

	res, err := persistence.UpdateSoftDeleteStateById(&ctx, false, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
