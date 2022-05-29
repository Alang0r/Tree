package lib

import (
	// "crypto/md5"
	// "encoding/hex"

	// "encoding/json"
	// "time"

	// "Tree/lib/log"

	// "github.com/streadway/amqp"
)

// type Header struct {
// 	// //Header ReqHeader//TODO: ADD crypto
// 	// RabbitChannel *amqp.Channel
// 	RequestID            string
// 	RequestName          string
// 	// To            string //destination service
// 	// Request       interface{}
// }

// //Send post request to RabbitMQ queue
// func (req *Header) Send() {
// 	logger := log.Logger{}
// 	logger.Init()

// 	//создаем reqId
// 	req.RequestId = GetMD5Hash(req.RequestName + time.Now().String())

// 	type RMQRequest struct {
// 		ID   string
// 		Name string
// 		To   string
// 		Req  interface{}
// 	}

// 	r := RMQRequest{
// 		ID: req.RequestId,
// 		Name: req.RequestName,
// 		To: req.To,
// 		Req: req.Request,
// 	}

// 	preMessage, err := json.Marshal(r)
// 	if err != nil {
// 		logger.Fatal(err.Error())
// 	}

// 	// r := RMQRequest{
// 	// 	ID: req.id,
// 	// 	Name: req.Name,
// 	// 	To: req.To,
// 	// 	Req: preMessage,
// 	// }
// 	// v, _ := json.Marshal(r)

// 	message := amqp.Publishing{
// 		ContentType: "application/json",
// 		Body:        preMessage,
// 	}

// 	err = req.RabbitChannel.Publish(
// 		"",
// 		req.To+"-queue",
// 		false,
// 		false,
// 		message,
// 	)
// 	if err != nil {
// 		logger.Fatal(err.Error())
// 	}

// }

// func GetMD5Hash(text string) string {
// 	hash := md5.Sum([]byte(text))
// 	return hex.EncodeToString(hash[:])
// }
