package api

import (
	"Tree/informer/models"
)

type ReqPersonCreate struct {
	Header
	Person models.Person
}

type RplPersonCreate struct {
	err error
}

func (req *ReqPersonCreate) GetInfo() (info map[string]string) {
	info = make(map[string]string)
	info["RequestName"] = "Informer/person/create"
	info["TimeOut"] = "5"
	info["Service"] = GetServiceName()
	return info
}

func (req *ReqPersonCreate) Run() error {
	return nil
}

func (req *ReqPersonCreate) Prepare(params map[string]interface{}) error {

	return nil
}
