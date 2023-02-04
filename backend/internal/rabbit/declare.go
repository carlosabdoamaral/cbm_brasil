package rabbit

import (
	"log"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
)

func DeclareQueue() {
	queue, err := common.RabbitChannel.QueueDeclare(
		common.RabbitQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	common.RabbitQueue = &queue
}
