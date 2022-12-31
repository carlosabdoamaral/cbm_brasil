package rabbitservice

import "github.com/carlosabdoamaral/cbm_brasil/backend/common"

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
		common.LogError(err.Error())
		return
	}

	common.RabbitQueue = &queue
}
