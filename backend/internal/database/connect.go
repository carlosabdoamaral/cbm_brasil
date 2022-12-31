package database

import (
	"database/sql"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(common.DB_DRIVER, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", common.DB_USER, common.DB_PASS, common.DB_NAME))
	if err != nil {
		return nil, err
	}

	common.Database = db
	return db, nil
}
