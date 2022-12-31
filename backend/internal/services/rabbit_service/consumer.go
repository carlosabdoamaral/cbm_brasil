package rabbitservice

import (
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
)

func Consumer() {
	msgs, err := common.RabbitChannel.Consume(
		common.RabbitQueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		common.LogError(err.Error())
		return
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			logMessage := fmt.Sprintf("New message received at %s", common.RabbitQueueName)
			common.LogInfo(logMessage)

			switch msg.Type {
			case "SOME":
				common.LogInfo("Some type identified")
			}
		}
	}()

	<-forever
}
