package test

import (
	api "Tree/informer/api"
	lib "Tree/lib/service"
)

func TestPersonCreate() {
	srv := lib.Service{}
	srv.SetName("Test")

	srv.Configure()

	req := api.ReqPersonCreate{
	}
	req.RabbitChannel = srv.RabbitChannel
	req.Person.FirstName = "Alex"
	
	req.Send()

	
}
