package persistence

import (
	"context"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
	pb "github.com/carlosabdoamaral/cbm_brasil/backend/protodefs/gen/proto"
)

func CreateAccount(ctx *context.Context, req *pb.CreateAccount) (*pb.AccountDetails, error) {
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
	query := fmt.Sprintf(`
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
	WHERE account_tb.id = %d
	`, id)

	db := common.Database
	rows, err := db.Query(query)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	res := &responses.AccountDetails{}
	err = persistence.ScanAccountDetails(rows, res)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return responses.NewAccountDetailsModelFromJSONToProto(res), nil
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

	res := &responses.AccountDetails{}
	err = persistence.ScanAccountDetails(rows, res)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return responses.NewAccountDetailsModelFromJSONToProto(res), nil
}

func EditAccountById(ctx *context.Context, req *pb.EditAccount) (*pb.AccountDetails, error) {
	updateAccountTbQuery := `
	UPDATE account_tb
	SET
		full_name = $1,
		email = $2,
		cpf = $3,
		birth_date = $4,
		passwd = $5
	WHERE
		account_tb.id = $6;
	`

	updateAccountLocationTbQuery := `
	UPDATE account_location_tb
	SET
		cep = $1,
		country = $2,
		state = $3,
		city = $4,
		neighborhood = $5,
		street = $6,
		place_number = $7,
		complement = $8
	WHERE account_location_tb.id_account = $9;
	`

	db := common.Database
	reqAsJSON := responses.NewEditAccountModelFromProtoToJSON(req)
	_, err := db.Exec(
		updateAccountTbQuery,
		reqAsJSON.FullName,
		reqAsJSON.Email,
		reqAsJSON.Cpf,
		reqAsJSON.BirthDate,
		reqAsJSON.Password,
		reqAsJSON.Id,
	)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	_, err = db.Exec(
		updateAccountLocationTbQuery,
		reqAsJSON.Location.CEP,
		reqAsJSON.Location.Country,
		reqAsJSON.Location.State,
		reqAsJSON.Location.City,
		reqAsJSON.Location.Neighborhood,
		reqAsJSON.Location.Street,
		reqAsJSON.Location.PlaceNumber,
		reqAsJSON.Location.Complement,
		reqAsJSON.Id,
	)
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	accountDetails, err := GetAccountDetailsById(ctx, req.GetId())
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return accountDetails, nil
}

func UpdateSoftDeleteStateById(ctx *context.Context, newState bool, req *pb.Id) (*pb.StatusResponse, error) {
	query := `UPDATE account_tb SET soft_deleted = $1 WHERE id = $2`

	db := common.Database
	_, err := db.Exec(query, newState, req.GetId())
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	return &pb.StatusResponse{
		Message: "Account soft deleted successfully updated",
	}, nil
}
