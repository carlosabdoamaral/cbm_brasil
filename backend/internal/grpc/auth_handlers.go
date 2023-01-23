package grpc

import (
	"context"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	persistence "github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence/auth"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func (s *AuthServer) SignInByEmail(ctx context.Context, req *pb.SignInByEmailRequest) (*pb.AccountDetails, error) {
	common.LogInfo("[GRPC] SignInByEmail")

	accountDetails, err := persistence.SignInByEmail(&ctx, req)
	if err != nil {
		return nil, err
	}

	return accountDetails, nil
}

func (s *AuthServer) SignInByCPF(ctx context.Context, req *pb.SignInByCPFRequest) (*pb.AccountDetails, error) {
	common.LogInfo("[GRPC] SignInByCPF")

	accountDetails, err := persistence.SignInByCPF(&ctx, req)
	if err != nil {
		return nil, err
	}

	return accountDetails, nil
}
