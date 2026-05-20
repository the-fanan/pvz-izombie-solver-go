package simulator

import (
	"github.com/google/uuid"
)

type Plant interface {
	ID() uuid.UUID
	Copy() Plant
	Name() PlantName
}

func NewPlant(name PlantName) Plant {
	return &plant{
		id:                       uuid.New(),
		name:                     name,
	}
}

type PlantName int
const (
	PlantNameUnknown PlantName = iota
	PlantNameBrain
	PlantNameSpikeWeed
	PlantNamePeaShooter
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

func (p *plant) Copy() Plant {
	return &plant{
		id:                       p.id,
		name:                     p.name,
	}
}
