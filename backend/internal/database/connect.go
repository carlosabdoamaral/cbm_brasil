package database

import (
	"database/sql"
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB, err error) {
	db, err = sql.Open(common.DB_DRIVER, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", common.DB_USER, common.DB_PASS, common.DB_NAME))
	if err != nil {
		common.LogError(err.Error())
		return nil, err
	}

	common.Database = db
	return
}
