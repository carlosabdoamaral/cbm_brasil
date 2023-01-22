package rabbit

import (
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
)

func StartConsuming() {
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
		fmt.Println("[*] Waiting for new messages")

		for m := range msgs {
			// Ex.:
			// if m.Type == common.NEW_ACCOUNT_REQUEST_QUEUE_TYPE {
			// 	NewAccountQueueHandler()
			// }

			fmt.Println(string(m.Body))
		}
	}()

	<-forever
}
