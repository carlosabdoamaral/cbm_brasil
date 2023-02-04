package rabbit

import (
	"strings"
	"time"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/streadway/amqp"
)

func PublishMessage(body []byte, messageType string) error {
	err := common.RabbitChannel.Publish(
		"",
		common.RabbitQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Timestamp:   time.Now(),
			Type:        strings.ToUpper(messageType),
			Body:        body,
		},
	)

	if err != nil {
		common.LogError(err.Error())
		return err
	}

	return nil
}
