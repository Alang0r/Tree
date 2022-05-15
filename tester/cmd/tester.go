package main

import (
	api "Tree/informer/api"
	lib "Tree/lib/service"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	srv := lib.Service{}
	srv.SetName("Tester")

	srv.Configure()
	wg.Add(1)
	go srv.Start()

	req := api.ReqPersonCreate{}
	req.RabbitChannel = srv.RabbitChannel
	req.Person.FirstName = "Alex"
	req.Send()
	wg.Wait()
}
