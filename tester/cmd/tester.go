package main

import lib "Tree/lib/service"

func main() {

	srv := lib.Service{}
	srv.SetName("Tester")

	srv.Configure()
	srv.Start()

	req := lib.CrossServiceRequest{
		RabbitChannel: srv.RabbitChannel,
		Request: "huy:pizda",
		To: "Informer",
	}
	req.Send()
}
