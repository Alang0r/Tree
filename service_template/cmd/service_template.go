package main

import (
	lib "Tree/lib/service"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	srv := lib.Service{}
	srv.SetName("Service_template")

	srv.Configure()
	wg.Add(1)
	go srv.Start()
	wg.Wait()
}
