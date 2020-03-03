package kittiesbundle

import uuid "github.com/satori/go.uuid"

//Kitty struct
type Kitty struct {
	ID        uuid.UUID
	Name      string
	Breed     string
	BirthDate string
	Errors    map[string]string
}

//NewKitty create a new kitty
func NewKitty(id string, name string, breed string, birthDate string) *Kitty {
	uuid, _ := uuid.FromString(id)
	return &Kitty{
		ID:        uuid,
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
