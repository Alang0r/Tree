package main

import (
	"Tree/lib/service"
)

//main 
func main() {
	srv := lib.Service{}
	srv.SetName("Informer")
	srv.Start()
}