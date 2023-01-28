package rabbit

import (
	"context"
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
			if m.Type == common.CREATE_OCCURRENCE_RABBIT_MSG_TYPE {
				ConsumeCreateOccurrenceRequest(context.Background(), m.Body)
			}

			if m.Type == common.ACCEPT_OCCURRENCE_RABBIT_MSG_TYPE {
				ConsumeAcceptOccurrenceRequest(context.Background(), m.Body)
			}
		}
	}()

	<-forever
}
