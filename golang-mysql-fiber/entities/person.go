package entities

import (
	"gorm.io/gorm"
)

type Person struct {
	ID        uint    `gorm:"primary key; autoIncrement" json:"id"`
	FirstName *string `json:"firstName,omitempty" gorm:"column:firstName"`
	LastName  *string `json:"lastName,omitempty" gorm:"column:lastName"`
	BirthDate *string `json:"birthDate,omitempty" gorm:"column:birthDate"`
}

type Tabler interface {
	TableName() string 
}

func (Person) TableName() string {
	return "person"
}

func MigratePerson(db *gorm.DB) error {
	err := db.AutoMigrate(&Person{})
	return err
}
