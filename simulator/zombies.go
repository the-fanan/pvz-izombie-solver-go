package simulator

import (
	"github.com/google/uuid"
)

type Zombie interface {
Name() ZombieName
	ID() uuid.UUID
	Copy() Zombie
	SunCost() int
	GetCurrentPosition() (CellID, int, int, int)
	Move(frame Frame)
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
	ZombieConeHead
	ZombieBucketHead

	zombieCount
)

func (zn ZombieName) SunCost() int {
	return mapZombieToSunCost()[zn]
}

func (zn ZombieName) WalkingSpeed() int {
	return mapZombieToWalkingSpeed()[zn]
}

func (zn ZombieName) GetNextPositionFunc() GetNextPositionFunc {
	return mapZombieToGetNextPositionFunc()[zn]
}

func mapZombieToSunCost() [zombieCount]int{
	return [...]int{
		ZombieUnknown: -1,
		ZombieImp: 50,
		ZombieConeHead: 75,
		ZombieBucketHead: 125,
	}
}

func mapZombieToWalkingSpeed() [zombieCount]int{
	return [...]int{
		ZombieUnknown: -1,
		ZombieImp: 3,
		ZombieConeHead: 3,
		ZombieBucketHead: 3,
	}
}

type GetNextPositionFunc func(frame Frame, zombie *zombie, speed float64) (CellID, int, int, int)

func defaultNextPositionFunc(frame Frame, zombie *zombie, speed float64) (CellID, int, int, int) {
	currentCell, currentSubmatrixRow, currentSubMatrixCol, currentAltitude := zombie.getStartingPosition()
	startingPoint := GetHorizontalDistanceFromCellColAndSubCellCol(currentCell.Position().Col, currentSubMatrixCol, zombie.resolution)

	// we use minus because we are going from right to left
	endPoint := startingPoint - int(speed * float64(zombie.movementFrame))
	newCol, newSubmatrixCol := GetCellColAndSubCellColFromHorizontalDistance(endPoint, zombie.resolution)
	newCellID := GetCellID(currentCell.Position().Row, newCol)

	return newCellID, currentSubmatrixRow, newSubmatrixCol, currentAltitude
}

func mapZombieToGetNextPositionFunc() [zombieCount]GetNextPositionFunc {
	return [...]GetNextPositionFunc{
		ZombieUnknown: func(frame Frame, zombie *zombie, speed float64) (CellID, int, int, int){
			return CellOutOfBounds, -1, -1, -1
		},
		ZombieImp: defaultNextPositionFunc,
		ZombieConeHead: defaultNextPositionFunc,
		ZombieBucketHead: defaultNextPositionFunc,
	}
}

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
	movementFrame int
	isFrozen bool
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
		movementFrame:         z.movementFrame,
	}
}

func (z *zombie) SunCost() int {
	return z.name.SunCost()
}

func (z *zombie) GetCurrentPosition() (CellID, int, int, int) {
	return z.cellID, z.subMatrixRow, z.subMatrixCol, z.subMatrixAltitude
}

func (z *zombie) Move(frame Frame) {
	// it's important that zombie wide modifications to speed are handled here
	// it reduces number of places to make updates if new modification requirements come up
	speed := float64(z.name.WalkingSpeed())
	if z.isFrozen {
		speed = speed / 2
	}

	relativeSpeed := speed / float64(z.resolution)
	// we must increment frame before getting next position
	z.movementFrame++
	c, sr, sc, sa := z.getNextPosition(frame, relativeSpeed)
	z.cellID = c
	z.subMatrixRow = sr
	z.subMatrixCol = sc
	z.subMatrixAltitude = sa

}

func (z *zombie) getStartingPosition() (CellID, int, int, int) {
	return z.startingCellID, z.startingSubMatrixRow, z.startingSubMatrixCol, z.startingSubMatrixAltitude
}

func (z *zombie) getNextPosition(frame Frame, speed float64) (CellID, int, int, int) {
	npf := z.name.GetNextPositionFunc()

	return npf(frame, z, speed)
}
