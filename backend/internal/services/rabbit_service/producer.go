package rabbitservice

import (
	"strings"
	"time"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/streadway/amqp"
)

func SendMessage(body []byte, t string) error {
	err := common.RabbitChannel.Publish(
		"",
		common.RabbitQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Timestamp:   time.Now(),
			Type:        strings.ToUpper(t),
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
