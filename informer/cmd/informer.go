package main

import (
	lib "Tree/lib/service"
)

//main
func main() {
	srv := lib.Service{}
	srv.SetName("Informer")

	srv.Configure()
	srv.Start()

}
