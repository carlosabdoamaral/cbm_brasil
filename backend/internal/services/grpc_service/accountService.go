package grpcservice

import (
	"context"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	AccountDb "github.com/carlosabdoamaral/cbm_brasil/backend/internal/database/account"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *AccountServer) GetAll(ctx context.Context, req *pb.Empty) (res *pb.GetAllAccountsResponse, err error) {
	common.LogInfo("[GRPC] GetAll")
	res, err = AccountDb.GetAll(ctx)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return &pb.GetAllAccountsResponse{
		Accounts: res.GetAccounts(),
	}, nil
}

func (s *AccountServer) PrivateDetails(ctx context.Context, req *pb.Id) (*pb.AccountPrivateInfos, error) {
	common.LogInfo("[GRPC] PrivateDetails")

	res, err := AccountDb.GetPrivateInfosById(ctx, req)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return res, nil
}

func (s *AccountServer) PublicDetails(ctx context.Context, req *pb.Id) (*pb.AccountPublicInfos, error) {
	common.LogInfo("[GRPC] PublicDetails")

	res, err := AccountDb.GetPublicInfosById(ctx, req)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return res, nil
}
