package kittiesbundle

import (
	"net/http"

	"github.com/carantes/kitties-api/app/core"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// KittiesRepositoryInterface define contract for the repository functions
type KittiesRepositoryInterface interface {
	FindAll() ([]Kitty, error)
	Delete(uuid.UUID) error
	Insert(*Kitty) error
	Get(uuid.UUID) (Kitty, error)
}

// KittiesController struct
type KittiesController struct {
	core.Controller
	repo KittiesRepositoryInterface
}

// NewKittiesController instance
func NewKittiesController(repo KittiesRepositoryInterface) *KittiesController {
	return &KittiesController{
		Controller: core.Controller{},
		repo:       repo,
	}
}

// Index return all kities
func (c *KittiesController) Index(w http.ResponseWriter, r *http.Request) {
	k, err := c.repo.FindAll()
	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}

// Get return one kitty by ID
func (c *KittiesController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uuid, err := uuid.FromString(vars["id"])
	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	k, err := c.repo.Get(uuid)
	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}

// Create a kitty
func (c *KittiesController) Create(w http.ResponseWriter, r *http.Request) {
	var k Kitty

	err := c.GetContent(&k, r)
	if c.HandleError(err, w) {
		return
	}

	if !k.Validate() {
		c.SendJSON(w, k.Errors, http.StatusBadRequest)
		return
	}

	err = c.repo.Insert(&k)
	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusCreated)
}

// Delete a kitty
func (c *KittiesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uuid, err := uuid.FromString(vars["id"])
	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	// TODO: reutrn the type of error to send 404 or 500
	err = c.repo.Delete(uuid)
	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, nil, http.StatusNoContent)
}
