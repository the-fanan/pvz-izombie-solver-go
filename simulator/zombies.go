package simulator

import "github.com/google/uuid"

type Zombie interface {
Name() ZombieName
	ID() uuid.UUID
	Copy() Zombie
	ExactCopy() Zombie
	SunCost() int
	SetResolution(resolution int)
}


type ZombieName int
const (
	ZombieUnknown ZombieName = iota
	ZombieImp

	zombieCount
)

func (zn ZombieName) SunCost() int {
	return mapZombieToSunCost()[zn]
}

func (zn ZombieName) WalkingSpeed() int {
	return mapZombieToWalkingSpeed()[zn]
}

func mapZombieToSunCost() [zombieCount]int{
	return [...]int{
		ZombieUnknown: -1,
		ZombieImp: 50,
	}
}

func mapZombieToWalkingSpeed() [zombieCount]int{
	return [...]int{
		ZombieUnknown: -1,
		ZombieImp: 1,
	}
}

type zombie struct {
	id                    uuid.UUID
	name                  ZombieName
	resolution int
}

func (z *zombie) Name() ZombieName {
	return z.name
}

func (z *zombie) ID() uuid.UUID {
	return z.id
}

func (z *zombie) copy() *zombie {
	return &zombie{
		id:                    uuid.New(),
		name:                  z.name,
		resolution:            z.resolution,
	}
}

func (z *zombie) Copy() Zombie {
	newZombie := z.copy()
	newZombie.id = uuid.New()
	return newZombie
}

func (z *zombie) ExactCopy() Zombie {
	newZombie := z.copy()
	newZombie.id = z.id
	return newZombie
}

func (z *zombie) SunCost() int {
	return z.name.SunCost()
}

func (z *zombie) SetResolution(resolution int) {
	z.resolution = resolution
}
