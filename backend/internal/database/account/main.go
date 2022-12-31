package accountDatabase

import (
	"context"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func GetAll(ctx context.Context) (*pb.GetAllAccountsResponse, error) {
	common.LogInfo("[DB  ] GetAll")

	query := `SELECT id, name, email, created_at FROM account_tb;`
	rows, err := common.Database.QueryContext(ctx, query)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	res := &pb.GetAllAccountsResponse{}
	for rows.Next() {
		model := pb.AccountPublicInfos{}
		rows.Scan(&model.Id, &model.Name, &model.Email, &model.CreatedAt)
		res.Accounts = append(res.Accounts, &model)
	}

	return res, nil
}

func GetPrivateInfosById(ctx context.Context, id *pb.Id) (*pb.AccountPrivateInfos, error) {
	common.LogInfo("[DB  ] GetById * Private")

	// TODO: Retornar token e location
	query := fmt.Sprintf(`
	SELECT
		account_tb.id,
		account_tb.name,
		account_tb.email,
		account_tb.password,
		account_tb.is_deleted,
		account_tb.created_at,
		account_tb.updated_at
	FROM
		account_tb
	WHERE
    	account_tb.id=%d;
	`, id.GetId())

	rows, err := common.Database.Query(query)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	res := &pb.AccountPrivateInfos{}
	for rows.Next() {
		rows.Scan(
			&res.Id,
			&res.Name,
			&res.Email,
			&res.Password,
			&res.IsDeleted,
			&res.CreatedAt,
			&res.UpdatedAt,
		)
	}

	return res, nil
}

func GetPublicInfosById(ctx context.Context, id *pb.Id) (*pb.AccountPublicInfos, error) {
	common.LogInfo("[DB  ] GetById * Public")

	query := fmt.Sprintf(`SELECT id, name, email, created_at FROM account_tb WHERE account_tb.id=%d;`, id.GetId())
	rows, err := common.Database.Query(query)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	res := &pb.AccountPublicInfos{}
	for rows.Next() {
		rows.Scan(
			&res.Id,
			&res.Name,
			&res.Email,
			&res.CreatedAt,
		)
	}

	return res, nil
}
