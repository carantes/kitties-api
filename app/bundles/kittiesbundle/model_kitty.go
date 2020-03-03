package kittiesbundle

import (
	"github.com/carantes/kitties-api/app/core"
)

//Kitty struct
type Kitty struct {
	core.Base
	Name      string `json:"name" gorm:"column:name;size:128;not null;"`
	Breed     string `json:"breed" gorm:"column:breed;size:128;"`
	BirthDate string `json:"birth_date" gorm:"column:birth_date;size:128;"`
}

//NewKitty create a new kitty
func NewKitty(name string, breed string, birthDate string) *Kitty {
	return &Kitty{
		Name:      name,
		Breed:     breed,
		BirthDate: birthDate,
	}
}

//Validate a Kitty
func (k *Kitty) Validate() bool {
	k.Errors = make(map[string]string)

	if k.Name == "" {
		k.Errors["name"] = "name can not be empty"
	}

	if len(k.Errors) > 0 {
		return false
	}

	return true
}
