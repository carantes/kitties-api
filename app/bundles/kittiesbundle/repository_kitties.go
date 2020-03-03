package kittiesbundle

import (
	"log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// KittiesRepository manage access to database
type KittiesRepository struct {
	db *gorm.DB
}

// NewKittiesRepository return repository instance
func NewKittiesRepository(db *gorm.DB) *KittiesRepository {
	return &KittiesRepository{
		db: db,
	}
}

// FindAll implement KittiesRepositoryInterface
func (repo *KittiesRepository) FindAll() ([]Kitty, error) {
	var kitties []Kitty

	log.Println("Get all kitties")
	repo.db.Find(&kitties)
	return kitties, nil
}

// Get implement KittiesRepositoryInterface
func (repo *KittiesRepository) Get(id uuid.UUID) (Kitty, error) {
	var kitty Kitty

	log.Printf("Get kitty with id = %v", id)
	err := repo.db.First(&kitty, "id = ?", id).Error
	return kitty, err
}

// Insert implement KittiesRepositoryInterface
func (repo *KittiesRepository) Insert(k *Kitty) error {
	log.Printf("Insert new kitty with name = %v", k.Name)
	return repo.db.Create(k).Error
}

// Delete implement KittiesRepositoryInterface
func (repo *KittiesRepository) Delete(id uuid.UUID) error {
	log.Printf("Delete kitty with id = %v", id)
	return repo.db.Delete(&Kitty{}, "id = ?", id).Error
}
