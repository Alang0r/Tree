package api

import (
	"Tree/informer/models"
)

type reqPersonCreate struct {
	person models.Person
}

type rplPersonCreate struct {
}

func (req *reqPersonCreate) Execute() {
	//добавляем нового родственника в дерево
	//req.person.DateBirth = 

}
