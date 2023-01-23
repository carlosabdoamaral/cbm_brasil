package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func SignInByEmail(ctx *context.Context, req *pb.SignInByEmailRequest) (*pb.AccountDetails, error) {
	query := `
	SELECT
		account_tb.id,
		account_tb.full_name,
		account_tb.email,
		account_tb.cpf,
		account_tb.birth_date,
		account_tb.passwd,
		account_tb.two_factor_code,
		account_location_tb.id,
		account_location_tb.id_account,
		account_location_tb.cep,
		account_location_tb.country,
		account_location_tb.state,
		account_location_tb.city,
		account_location_tb.neighborhood,
		account_location_tb.street,
		account_location_tb.place_number,
		account_location_tb.complement,
		account_tb.created_at,
		account_tb.updated_at,
		account_tb.soft_deleted
	FROM account_tb
	INNER JOIN account_location_tb on account_tb.id = account_location_tb.id_account
	WHERE email = $1 AND passwd = $2
	LIMIT 1
	`

	db := common.Database
	rows, err := db.Query(query, req.GetEmail(), req.GetPassword())
	if err == sql.ErrNoRows {
		return nil, errors.New("no account founded with this credentials")
	}
	if err != nil {
		return nil, err
	}

	res := &responses.AccountDetailsJSON{}
	err = persistence.ScanAccountDetails(rows, res)
	if err != nil {
		return nil, err
	}

	return responses.NewAccountDetailsFromJSONToProto(res), nil
}
