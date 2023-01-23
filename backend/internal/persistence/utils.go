package persistence

import (
	"database/sql"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/responses"
)

func ScanAccountDetails(rows *sql.Rows, res *responses.AccountDetailsJSON) error {
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
			return err
		}
	}

	return nil
}
