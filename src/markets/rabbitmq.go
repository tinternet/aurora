package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

const (
	tickersExchangeName string = "tickers"
)

var (
	rabbitConnection *amqp.Connection
)

func init() {
	mqAddr := os.Getenv("RABBITMQ_ADDR")
	conn, err := amqp.Dial("amqp://guest:guest@" + mqAddr + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	rabbitConnection = conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publishTickers(channel <-chan *ticker) {
	ch, err := rabbitConnection.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		tickersExchangeName,
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare tickers exchange")

	for ticker := range channel {
		body, err := json.Marshal(&ticker)
		if err != nil {
			log.Printf("Cannot encode ticker: %s", err)
			continue
		}
		routingKey := fmt.Sprintf("ticker.%s.%s", strings.ToLower(ticker.Market), strings.ToLower(ticker.Symbol))
		err = ch.Publish(
			tickersExchangeName,
			routingKey,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		if err != nil {
			log.Printf("Cannot publish ticker: %s", err)
		}
	}
}

func consumeTickers() {
	ch, err := rabbitConnection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		tickersExchangeName, // name
		"fanout",            // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,              // queue name
		"*.*.*",             // routing key
		tickersExchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
