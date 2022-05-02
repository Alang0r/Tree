package models

const tablePerson = "Person"
type Person struct {
	id        int
	DateBirth int
	Datedeath int
	FirstName string
	Surname   string
	Lastname  string
	MotherId  int
	FatherId  int
}

func (Person) GetTableName() string {
	return tablePerson
}
