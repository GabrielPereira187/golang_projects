package main

import (
	"encoding/json"
	"fmt"
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


func main(){
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	msgs, err := channel.Consume(
		"testing", // queue
			"",        // consumer
			true,      // auto ack
			false,     // exclusive
			false,     // no local
			false,     // no wait
			nil,       //args
	)

	if err != nil {
		panic(err)
	}


	var payment Payment

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			body := msg.Body

			json.Unmarshal(body, &payment)

			payment.Consumed = true
			fmt.Println("Received Message:", payment)
		}
	}()

	fmt.Println("Processing messages...") 
	<- forever
}

	


