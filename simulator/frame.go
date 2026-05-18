package simulator

type Frame struct {
	CellResolution int
	RowToTrack                   int16
	Board                        [][]Cell
	Iteration                    int64
	LastZombieInsertionIteration int64
	ZombieInsertions             []ZombieInsertion
}

type ZombieInsertion struct {
	Iteration  int64
	ZombieName ZombieName
}

func (f *Frame) CountZombies() int {
	zombieCount := 0
	rows := int16(len(f.Board))
	cols := int16(len(f.Board[0]))

	for row := int16(0); row < rows; row++ {
		for col := int16(0); col < cols; col++ {
			zombieCount += f.Board[row][col].Zombies.Len()
		}
	}

	return zombieCount
}

func (f *Frame) SetCellPlant(plant Plant, row, col int16) {
	f.Board[row][col].Plants.PushBack(plant)
}

func (f *Frame) RemoveCellPlant(plant Plant, row, col int16) {
	if f.Board[row][col].Plants.Len() == 0 {
		return
	}
	element := f.Board[row][col].Plants.Front()

	for element != nil {
		otherPlant := element.Value.(Plant)
		if otherPlant.ID() == plant.ID() {
			f.Board[row][col].Plants.Remove(element)
			break
		}

		element = element.Next()
	}
}

func (f *Frame) SetCellZombie(zombie Zombie, row, col int16) {
	f.Board[row][col].Zombies.PushBack(zombie)
	f.LastZombieInsertionIteration = f.Iteration
	f.ZombieInsertions = append(f.ZombieInsertions, ZombieInsertion{
		Iteration:  f.Iteration,
		ZombieName: zombie.Name(),
	})
}

func (f *Frame) RemoveCellZombie(zombie Zombie, row, col int16) {
	if f.Board[row][col].Zombies.Len() == 0 {
		return
	}
	element := f.Board[row][col].Zombies.Front()

	for element != nil {
		otherZombie := element.Value.(Zombie)
		if otherZombie.ID() == zombie.ID() {
			f.Board[row][col].Zombies.Remove(element)
			break
		}

		element = element.Next()
	}
}
