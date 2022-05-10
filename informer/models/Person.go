package models

const tablePerson = "person"

type Person struct {
	id        int    `gorm: "id"`
	DateBirth int    `gorm: "date_birth"`
	Datedeath int    `gorm: date_death"`
	FirstName string `gorm: "first_name"`
	Surname   string `gorm: "surname"`
	Lastname  string `gorm: "last_name"`
	MotherId  int    `gorm: "mother_id"`
	FatherId  int    `gorm: "father_id"`
}

func (Person) GetTableName() string {
	return tablePerson
}
