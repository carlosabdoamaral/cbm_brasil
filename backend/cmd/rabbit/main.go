package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
)

func main() {
	common.ReadEnvVariables()

	rabbit.Connect()
	rabbit.DeclareQueue()
	rabbit.StartConsuming()
}
