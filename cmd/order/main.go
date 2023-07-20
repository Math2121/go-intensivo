package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Math2121/go-first-step/internal/infra/database"
	usecase "github.com/Math2121/go-first-step/internal/useCase"
	"github.com/Math2121/go-first-step/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	caseUse := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitMqWorker(msgRabbitmqChannel, caseUse)
}

func rabbitMqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)

		fmt.Println("mensagem processada ", output)
	}
}
