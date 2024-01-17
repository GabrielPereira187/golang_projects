package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

type Payment struct {
	ID int
	Value float32
	AccountFrom string
	AccountTo string
	PaymentDate time.Time
	Consumed bool
}

const queueName = "testing"

func main() {
	fmt.Println("Rabbit Starting")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if  err != nil {
		panic(err)
	}

	fmt.Println("Rabbit connected")

	channel, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	payment := Payment{
		ID: rand.Intn(1000),
		Value: rand.Float32(),
		AccountFrom: "Gabriel",
		AccountTo: "Joao",
		PaymentDate: time.Now(),
		Consumed : false,
	}

	body, err := json.Marshal(payment)
	if err != nil {
		panic(err)
	}


	err = channel.Publish(
		"",
		"testing",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
}