package util

import (
	"Tree/lib/log"
	"encoding/json"

	"github.com/streadway/amqp"
)

type InnerRequest struct {
	Header
	Request interface{}
	Channel *amqp.Channel
}

type Request interface{
	Name() string
}


func (req *InnerRequest) SendRequest() (resp string, err error) {
	logger := log.Logger{}
	logger.Init()

	req.SetID()

	//создаем очередь, в которой будем ожидать ответ
		_, err = req.Channel.QueueDeclare(
			req.RequestID,
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			logger.Fatal(err.Error())
		}

	//готовим сообщение для отправки
	preMessage, err := json.Marshal(req.Request)
	if err != nil {
		logger.Fatal(err.Error())
	}
	headers := make(map[string] interface{})

	headers["ReqID"] = req.RequestID
	headers["ReqName"] = req.RequestName
	headers["Recipient"] = req.Recipient

	message := amqp.Publishing{
		Headers: headers,
		ContentType: "application/json",
		Body:        preMessage,
	}

	//отправляем сообщение в очередь соответствующего сервиса
	err = req.Channel.Publish(
		"",
		req.Recipient + "-queue",
		false,
		false,
		message,
	)

	if err != nil {
		return "", err
	}

	//ждем ответ //TODO:до истечения указанного таймаута

	//Читаем сообщение и вызываем соответствующий запрос / возвращаем ошибку, если запрсоа нет
		msgs, err := req.Channel.Consume(
			req.RequestID,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			logger.Info(err.Error())
			return "", err
		}
		var response []byte
		for resp := range msgs {
				logger.Info("Response: " + string(resp.Body))
				response = resp.Body
			}

	return string(response), nil
}
