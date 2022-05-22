package request

type Request interface	{
	Execute() (resp Response, err error)
}

type Response interface {
	
}
type Header struct {
	RequestID   string
	RequestName string
}
// import (
// 	"Tree/lib/log"
// 	// lib "Tree/lib/service"
// 	"crypto/md5"
// 	"encoding/hex"
// 	"encoding/json"
// 	"time"

// 	"github.com/streadway/amqp"
// )

// type CrossServiceRequest struct {
// 	id        string
// 	Name      string
// 	Recipient string
// 	Request   interface{}
// 	Channel   *amqp.Channel
// }

// // type Header struct {
// // 	ReqId string
// // 	ReqName string
// // }

// //Send post request to RabbitMQ queue
// func (obj *CrossServiceRequest) Send() {
// 	logger := log.Logger{}
// 	logger.Init()

// 	type Request struct {
// 		Header Header
// 		Request interface {
// 		}
// 	}
// 	extReq := Req{
// 		// ID: GetMD5Hash(req.Name + time.Now().String()),
// 		Header.RequestID: 	GetMD5Hash(Header.RequestName + time.Now().String()),
// 	}
// 	//создаем reqId
// 	req.Header.ID = GetMD5Hash(req.Name + time.Now().String())

// 	preMessage, err := json.Marshal(r)
// 	if err != nil {
// 		logger.Fatal(err.Error())
// 	}
// 	message := amqp.Publishing{
// 		ContentType: "application/json",
// 		Body:        preMessage,
// 	}

// 	err = req.Channel.Publish(
// 		"",
// 		req.Recipient+"-queue",
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

// type Request struct {
// 	Header
// 	Req interface{}
// }
// func SendRequest(req Request, ch *amqp.Channel) (resp interface{}, err error) {
// 	logger := log.Logger{}
// 	logger.Init()

// 	req.RequestID = GetMD5Hash(req.RequestName + time.Now().String())

// 	preMessage, err := json.Marshal(req)
// 	if err != nil {
// 		logger.Fatal(err.Error())
// 	}


// 	message := amqp.Publishing{
// 		ContentType: "application/json",
// 		Body:        preMessage,
// 	}

// 	err = ch.Publish(
// 		"",
// 		req.Recipient + "-queue",
// 		false,
// 		false,
// 		message,
// 	)
// 	if err != nil {
// 		logger.Fatal(err.Error())
// 	}

// 	return nil, nil
// }