package test

import (
	lib "Tree/lib/service"
	"Tree/lib/util"
	"testing"
	 informer "Tree/informer/api"
)

func TestPersonCreate(t *testing.T) {
	srv := lib.Service{}
	srv.SetName("Test")

	srv.Configure()

	req := informer.ReqPersonCreate{}
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
