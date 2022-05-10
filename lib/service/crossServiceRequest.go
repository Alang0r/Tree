package lib

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type CrossServiceRequest struct {
	//Header ReqHeader//TODO: ADD crypto
	RabbitChannel *amqp.Channel
	Request       string //json with all parameters
	To            string //destination service
}


//Send post request to RabbitMQ queue
func (req *CrossServiceRequest) Send() {

	preMessage, err := json.Marshal(req.Request)

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(preMessage),
	}

	err = req.RabbitChannel.Publish(
		"",
		req.To + "-queue",
		false,
		false,
		message,
	)
	if err != nil {
		panic(err)
	}

}
