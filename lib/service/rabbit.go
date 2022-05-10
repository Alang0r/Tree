package lib

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Queue struct {
}

func (Queue) Init() {
	rabbitConfig := "amqp://guest:guest@" +
		"" +
		":5672"
	con, err := amqp.Dial(rabbitConfig)
	if err != nil {
		panic(err)
	}

	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"", //FIXME: add queue
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(q)

	err = ch.Publish(
		"",
		"", //FIXME add queue
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte("Ololololoo"),
		},
	)
	if err != nil {
		panic(err)
	}
}
