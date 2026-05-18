package simulator

import (
	"github.com/google/uuid"
)

type Plant interface {
	ID() uuid.UUID
	Copy() Plant
	ExactCopy() Plant
	Name() PlantName
}

type PlantName int
const (
	PlantNameUnknown PlantName = iota
	PlantNameBrain
)

type plant struct {
	id uuid.UUID
	name PlantName
}

func (p *plant) Name() PlantName {
	return p.name
}

func (p *plant) ID() uuid.UUID {
	return p.id
}

// to be used when initializing a frame
func (p *plant) Copy() Plant {
	newPlant := p.copy()
	newPlant.id = uuid.New()
	return newPlant
}

// to be used for generating frame continuation
func (p *plant) ExactCopy() Plant {
	newPlant := p.copy()
	newPlant.id = p.id
	return newPlant
}

func (p *plant) copy() *plant {
	return &plant{
		id:                       uuid.New(),
		name:                     p.name,
	}
}
