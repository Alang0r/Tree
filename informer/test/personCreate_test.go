package test

import (
	"Tree/informer/api"
	lib "Tree/lib/service"
	"fmt"

	"Tree/lib/util"
	"testing"
)

func TestPersonCreate(t *testing.T) {
	srv := lib.Service{}
	srv.SetName("Test")

	srv.Configure()

	req := api.ReqPersonCreate{}
	reqData := make(map[string]interface{})
	reqData["Id"] = 0
	reqData["DateBirth"] = 1994
	reqData["Datedeath"] = 0
	reqData["FirstName"] = "Aleksandr"
	reqData["Surname"] = "Andreevich"
	reqData["Lastname"] = "Gorodilov"
	reqData["MotherId"] = 0
	reqData["FatherId"] = 0
	lib.FillStruct(reqData, &req.Person)
	fmt.Println(req.Person)
	newInnerRequest := util.InnerRequest{}
	info := req.GetInfo()

	newInnerRequest.Channel = srv.RabbitChannel
	newInnerRequest.Recipient = info["Service"]
	newInnerRequest.RequestName = info["RequestName"]
	newInnerRequest.Request = req
	resp, err := newInnerRequest.SendRequest()
	if err != nil {
		srv.Log.Fatalf("%s Returned error: %s", info["Queue"], err.Error())
	}
	srv.Log.Info(resp)

	// req := api.ReqPersonCreate{
	// }
	// req.RabbitChannel = srv.RabbitChannel
	// req.To = "Informer"
	// req.Person.FirstName = "Alex"

	// req.Send()

}
