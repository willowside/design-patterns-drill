package pets

import "go-designpattern/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description",
		Lifespan:    0,
	}
	return &pet
}
