package pets

import "errors"

type PetInterface interface {
	SetSepcies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(s int) *Pet
	SetMaxWeight(s int) *Pet
	SetWeight(s int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(s int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(s int) *Pet
	SetAgeEstimated(s bool) *Pet
	Build() (*Pet, error)
}

// small factory of the Builder
func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetSepcies(s string) *Pet {
	p.Species = s
	return p
}
func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}
func (p *Pet) SetMinWeight(s int) *Pet {
	p.MinWeight = s
	return p
}
func (p *Pet) SetMaxWeight(s int) *Pet {
	p.MaxWeight = s
	return p
}
func (p *Pet) SetWeight(s int) *Pet {
	p.Weight = s
	return p
}
func (p *Pet) SetDescription(s string) *Pet {
	p.Description = s
	return p
}
func (p *Pet) SetLifeSpan(s int) *Pet {
	p.LifeSpan = s
	return p
}
func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}
func (p *Pet) SetColor(s string) *Pet {
	p.Color = s
	return p
}
func (p *Pet) SetAge(s int) *Pet {
	p.Age = s
	return p
}
func (p *Pet) SetAgeEstimated(s bool) *Pet {
	p.AgeEstimated = s
	return p
}

func (p *Pet) Build() (*Pet, error) {
	// validity checks
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}
	p.Average = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
