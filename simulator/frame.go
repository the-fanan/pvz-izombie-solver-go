package simulator

import (
	"container/list"
)

type Frame interface {
	GetCellResolution() int
	CountZombies() int
	SetCellPlant(plant Plant, row, col int)
	RemoveCellPlant(plant Plant, row, col int)
	SetCellZombie(zombie Zombie, row, col int)
	RemoveCellZombie(zombie Zombie, row, col int)
	CellHasPlant(cellID CellID) bool
}

func NewFrame(rows, cols, cellResolution, rowToSimulateOn, sun int) Frame {
	// ensure rows is at most 5
	NUM_ROWS := 5
	if rows < NUM_ROWS {
		NUM_ROWS = rows
	}
	// ensure columns is at most 7
	NUM_COLS := 7
	if cols < NUM_COLS {
		NUM_COLS = cols
	}
	board := [][]Cell{}

	for row := 0; row < NUM_ROWS; row++ {
		board = append(board, []Cell{})
		for col := 0; col < NUM_COLS; col++ {
			cellID := GetCellID(row, col)
			if cellID == CellOutOfBounds {
				continue
			}
			board[row] = append(board[row], Cell{
				CellID:            cellID,
				Position: CellPosition{Row: row, Col: col},
				Plants:            list.New(),
				Zombies:           list.New(),
			})
		}
	}
	return &frame {
		CellResolution: cellResolution,
		RowToSimulateOn: rowToSimulateOn,
		Board: board,
		Sun: sun,
	}
}

type frame struct {
	CellResolution int
	RowToSimulateOn                   int
	Board                        [][]Cell
	Iteration                    int64
	LastZombieInsertionIteration int64
	ZombieInsertions             []ZombieInsertion
	Sun int
}

type ZombieInsertion struct {
	Iteration  int64
	ZombieName ZombieName
}

func (f *frame) CountZombies() int {
	zombieCount := 0
	rows := int(len(f.Board))
	cols := int(len(f.Board[0]))

	for row := int(0); row < rows; row++ {
		for col := int(0); col < cols; col++ {
			zombieCount += f.Board[row][col].Zombies.Len()
		}
	}

	return zombieCount
}

func (f *frame) SetCellPlant(plant Plant, row, col int) {
	f.Board[row][col].Plants.PushBack(plant)
}

func (f *frame) RemoveCellPlant(plant Plant, row, col int) {
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

func (f *frame) SetCellZombie(zombie Zombie, row, col int) {
	f.Board[row][col].Zombies.PushBack(zombie)
	f.LastZombieInsertionIteration = f.Iteration
	f.ZombieInsertions = append(f.ZombieInsertions, ZombieInsertion{
		Iteration:  f.Iteration,
		ZombieName: zombie.Name(),
	})
}

func (f *frame) RemoveCellZombie(zombie Zombie, row, col int) {
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

func (f *frame) GetCellResolution() int {
	return f.CellResolution
}

func (f *frame) CellHasPlant(cellID CellID) (bool) {
	position := cellID.Position()
	cell := f.Board[position.Row][position.Col]
	plantElement := cell.Plants.Front()
	if plantElement == nil {
		return false
	}

	return true
}
