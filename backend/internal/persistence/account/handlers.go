package persistence

import (
	"context"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func CreateAccount(ctx *context.Context, req *pb.NewAccountRequest) (*pb.AccountDetails, error) {
	query := `
	DO $$
		DECLARE lastid integer;
		BEGIN
			INSERT INTO account_tb(full_name, email, cpf, birth_date)
			VALUES ('TEMPLATE', 'TEMPLATE', 'TEMPLATE', NOW())
			RETURNING id INTO lastid;

			INSERT INTO account_location_tb(id_account, cep, country, state, city, neighborhood, street, place_number, complement)
			VALUES (lastid, '00000000', 'BRAZIL', 'TEMPLATE', 'TEMPLATE', 'TEMPLATE', 'TEMPLATE', 999, 'TEMPLATE');
	END $$;
	`

	db := common.Database

	_, err := db.Exec(query)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	accountDetails, err := GetLastAccountDetails(ctx)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return accountDetails, nil
}

func GetAccountDetailsById(ctx *context.Context, id int64) (*pb.AccountDetails, error) {
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
	WHERE account_tb.id = ?
	`

	db := common.Database
	rows, err := db.Query(query, id)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	res := &pb.AccountDetails{}
	for rows.Next() {
		err := rows.Scan(
			&res.Id,
			&res.FullName,
			&res.Email,
			&res.Cpf,
			&res.BirthDate,
			&res.Passwd,
			&res.TwoFactorCode,
			&res.Location.Id,
			&res.Location.IdAccount,
			&res.Location.Cep,
			&res.Location.Country,
			&res.Location.State,
			&res.Location.City,
			&res.Location.Neighborhood,
			&res.Location.Street,
			&res.Location.PlaceNumber,
			&res.Location.Complement,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.SoftDeleted,
		)

		if err != nil {
			common.LogError(err.Error())
			return nil, err
		}
	}

	return res, nil
}

func GetLastAccountDetails(ctx *context.Context) (*pb.AccountDetails, error) {
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
	WHERE account_tb.id = (SELECT id FROM account_tb ORDER BY id DESC LIMIT 1)
	`

	db := common.Database
	rows, err := db.Query(query)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	res := &responses.AccountDetailsRequestJSON{}
	for rows.Next() {
		err := rows.Scan(
			&res.Id,
			&res.FullName,
			&res.Email,
			&res.Cpf,
			&res.BirthDate,
			&res.Password,
			&res.TwoFactorCode,
			&res.Location.Id,
			&res.Location.IdAccount,
			&res.Location.CEP,
			&res.Location.Country,
			&res.Location.State,
			&res.Location.City,
			&res.Location.Neighborhood,
			&res.Location.Street,
			&res.Location.PlaceNumber,
			&res.Location.Complement,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.SoftDeleted,
		)

		if err != nil {
			common.LogError(err.Error())
			return nil, err
		}
	}

	return responses.NewAccountDetailsFromJSONToProto(res), nil
}
