package api

import (
	"Tree/informer/models"
	lib "Tree/lib/service"
)

type reqPersonCreate struct {
	person models.Person
}

type rplPersonCreate struct {
	err error
}

func (req *reqPersonCreate) Execute() {
	outerReq := lib.CrossServiceRequest{}
	outerReq.Send()
}
