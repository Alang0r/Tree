package main

import (
	"Informer/api"
	"Tree/informer/api"
	lib "Tree/lib/service"
)

//main
func main() {
	srv := lib.Service{}
	srv.SetName("Informer")

	srv.Configure()
	srv.RegisterRequest("ReqPersonCreate", api.ReqPersonCreate)
	srv.Start()
}
