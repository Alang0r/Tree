package api

import (
	"Tree/informer/models"
	s "Tree/lib/service"
)
const(
	RequestName = "Informer/person/create"
)

type ReqPersonCreate struct {
	s.Header
	Person models.Person
}

type RplPersonCreate struct {
	err error
}

func (req *ReqPersonCreate) Init() {
	req.Name = RequestName
}

func (req *ReqPersonCreate) Execute() {
	
}
