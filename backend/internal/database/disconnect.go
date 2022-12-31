package database

import "github.com/carlosabdoamaral/cbm_brasil/backend/common"

func Disconnect() {
	common.Database.Close()
}
