package kittiesbundle

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// KittiesRepository manage access to database
type KittiesRepository struct {
	kitties []Kitty
}

// NewKittiesRepository return repository instance
func NewKittiesRepository() *KittiesRepository {
	return &KittiesRepository{
		kitties: []Kitty{
			*NewKitty("904bd6f4-8a13-4bcf-a77b-6f7a3c1b5dfd", "Gaspart", "British", "2016-07-05"),
			*NewKitty("c412d936-a69b-4c11-9d04-327e07d57a4f", "Marcel", "European", "2014-05-02"),
		},
	}
}

// FindAll implement KittiesRepositoryInterface
func (repo *KittiesRepository) FindAll() ([]Kitty, error) {
	return repo.kitties, nil
}

// Insert implement KittiesRepositoryInterface
func (repo *KittiesRepository) Insert(k *Kitty) error {
	k.ID = uuid.NewV4()
	repo.kitties = append(repo.kitties, *k)
	return nil
}

// Delete implement KittiesRepositoryInterface
func (repo *KittiesRepository) Delete(id uuid.UUID) error {

	for i, kitty := range repo.kitties {
		if kitty.ID == id {
			repo.kitties = append(repo.kitties[:i], repo.kitties[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Kitty with id '%v' not found", id)
}
