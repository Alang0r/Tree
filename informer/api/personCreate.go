package api

import (
	"Tree/informer/models"
	m "Tree/informer/models"
)

type reqPersonCreate struct {
	person models.Person
}

type rplPersonCreate struct {
	err error
}

func (req *reqPersonCreate) Execute() {
	//добавляем нового родственника в дерево
	newPerson := m.Person{
		Id:        req.person.Id,
		DateBirth: req.person.DateBirth,
		Datedeath: req.person.Datedeath,
		FirstName: req.person.FirstName,
		Surname: req.person.Surname,
		Lastname: req.person.Lastname,
		MotherId: req.person.MotherId,
		FatherId: req.person.FatherId,
	}

	//err := 


}
