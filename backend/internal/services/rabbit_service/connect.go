package rabbitservice

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/streadway/amqp"
)

func Connect() {
	conn, err := amqp.Dial(common.RABBIT_URL)
	if err != nil {
	}

	ch, err := conn.Channel()
	if err != nil {
	}

	common.RabbitConn = conn
	common.RabbitChannel = ch
}
