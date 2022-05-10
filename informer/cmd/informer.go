package main

import (
	"Tree/lib/service"
)

//main
func main() {
	srv := lib.Service{}
	srv.SetName("Informer")

	srv.Configure()
	srv.Start()
	// newPerson := models.Person{

	// }
	// srv.DB.Table(newPerson.GetTableName()).Select("*").First(&newPerson)
	// fmt.Println(newPerson)
}