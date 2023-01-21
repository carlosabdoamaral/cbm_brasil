package main

import (
	"fmt"
	"log"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
)

func main() {
	common.ReadEnvVariables()
	rabbit.Connect()

	DeclareQueue()
	ConsumeQueue()
}

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

func ConsumeQueue() {
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
		log.Println("Error")
	}

	forever := make(chan bool)
	go func() {
		fmt.Println("[*] Waiting for new messages")

		for m := range msgs {
			// Aqui terão switch case para redirecionar para o GRPC

			// Ex.:
			//// Account
			// switch m.Type {
			// case "NEWACCOUNT":
			//     NewAccountHandler(&m)
			// }

			fmt.Println(string(m.Body))
		}
	}()

	<-forever
}
