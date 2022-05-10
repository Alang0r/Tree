package lib

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Queue struct {
}

func (Queue) Init() {
	rabbitConfig := "amqp://guest:guest@"+
	""+
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
		"",//FIXME: add queue
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
		"",//FIXME add queue
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte("Ololololoo"),
		},
	)
	if err != nil {
		panic(err)
	}
}

//метод для отправки запроса в другой сервис через rabbitmq
//параметры для подключения (и подключение / канал?) у нас уже есть - из настроек нашего сервиса
type CrossServiceRequest struct {
	//Header ReqHeader//TODO: ADD crypto
	request string //json with all parameters


}

func (req *CrossServiceRequest) Send(ch *amqp.Channel) () {
	preMessage, err := json.Marshal(req.request)

	message := amqp.Publishing{
		ContentType: "application/json",
		Body: []byte(preMessage),
	}

	err = ch.Publish(
		"",
		"",//FIXME add queue
		false,
		false,
		message,
	)
	if err != nil {
		panic(err)
	}

}