package simulator

import "github.com/google/uuid"

type Zombie interface {
Name() ZombieName
	ID() uuid.UUID
	Copy() Zombie
	SunCost() int
	GetCurrentPosition() (CellID, int, int, int)
}

func NewZombie(name ZombieName, resolution int, startingCell CellID, startingSubmatrixRow, startingSubmatrixCol, startingSubmatrixAltitude int) Zombie {
	z := &zombie{
		id:                    uuid.New(),
		name:                  name,
		resolution:            resolution,
		cellID:                startingCell,
		subMatrixRow:          startingSubmatrixRow,
		subMatrixCol:          startingSubmatrixCol,
		subMatrixAltitude:     startingSubmatrixAltitude,
		startingCellID:        startingCell,
		startingSubMatrixRow:  startingSubmatrixRow,
		startingSubMatrixCol:  startingSubmatrixCol,
		startingSubMatrixAltitude: startingSubmatrixAltitude,
	}

	return z
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

type GetNextPosition func(frame Frame, zombie Zombie) (CellID, int, int, int)

type zombie struct {
	id                    uuid.UUID
	name                  ZombieName
	resolution int
	cellID CellID
	subMatrixRow int
	subMatrixCol int
	subMatrixAltitude int
	startingCellID CellID
	startingSubMatrixRow int
	startingSubMatrixCol int
	startingSubMatrixAltitude int
	startingFrame int
	movementFrame int // only increments when position is updated
}

func (z *zombie) Name() ZombieName {
	return z.name
}

func (z *zombie) ID() uuid.UUID {
	return z.id
}

func (z *zombie) Copy() Zombie {
	return &zombie{
		id:                    z.id,
		name:                  z.name,
		resolution:            z.resolution,
		cellID:                z.cellID,
		subMatrixRow:          z.subMatrixRow,
		subMatrixCol:          z.subMatrixCol,
		subMatrixAltitude:     z.subMatrixAltitude,
		startingCellID:        z.startingCellID,
		startingSubMatrixRow:  z.startingSubMatrixRow,
		startingSubMatrixCol:  z.startingSubMatrixCol,
		startingSubMatrixAltitude: z.startingSubMatrixAltitude,
		startingFrame:         z.startingFrame,
		movementFrame:         z.movementFrame,
	}
}

func (z *zombie) SunCost() int {
	return z.name.SunCost()
}

func (z *zombie) GetCurrentPosition() (CellID, int, int, int) {
	return z.cellID, z.subMatrixRow, z.subMatrixCol, z.subMatrixAltitude
}
