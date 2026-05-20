package simulator

import (
	"container/list"
)

type Cell struct {
	CellID CellID
	Position CellPosition
	Zombies *list.List
	Plants *list.List
}

type CellPosition struct {
	Row int
	Col int
}

type CellID int

const (
	CellOutOfBounds CellID = iota
	CellA0
	CellA1
	CellA2
	CellA3
	CellA4
	CellA5
	CellA6
	CellB0
	CellB1
	CellB2
	CellB3
	CellB4
	CellB5
	CellB6
	CellC0
	CellC1
	CellC2
	CellC3
	CellC4
	CellC5
	CellC6
	CellD0
	CellD1
	CellD2
	CellD3
	CellD4
	CellD5
	CellD6
	CellE0
	CellE1
	CellE2
	CellE3
	CellE4
	CellE5
	CellE6

	cellIDCount
)

func (c CellID) String() string {
	return mapCellIDToString()[c]
}

func (c CellID) Position() CellPosition {
	return mapCellIDToPosition()[c]
}

func mapCellIDToString() [cellIDCount]string {
	return [...]string{
		CellOutOfBounds: "Unknown",
		CellA0:      "A0",
		CellA1:      "A1",
		CellA2:      "A2",
		CellA3:      "A3",
		CellA4:      "A4",
		CellA5:      "A5",
		CellA6:      "A6",
		CellB0:      "B0",
		CellB1:      "B1",
		CellB2:      "B2",
		CellB3:      "B3",
		CellB4:      "B4",
		CellB5:      "B5",
		CellB6:      "B6",
		CellC0:      "C0",
		CellC1:      "C1",
		CellC2:      "C2",
		CellC3:      "C3",
		CellC4:      "C4",
		CellC5:      "C5",
		CellC6:      "C6",
		CellD0:      "D0",
		CellD1:      "D1",
		CellD2:      "D2",
		CellD3:      "D3",
		CellD4:      "D4",
		CellD5:      "D5",
		CellD6:      "D6",
		CellE0:      "E0",
		CellE1:      "E1",
		CellE2:      "E2",
		CellE3:      "E3",
		CellE4:      "E4",
		CellE5:      "E5",
		CellE6:      "E6",
	}
}

func mapCellIDToPosition() [cellIDCount]CellPosition {
	return [...]CellPosition{
		CellOutOfBounds: {Row: -1, Col: -1},
		CellA0:          {Row: 0, Col: 0},
		CellA1:          {Row: 0, Col: 1},
		CellA2:          {Row: 0, Col: 2},
		CellA3:          {Row: 0, Col: 3},
		CellA4:          {Row: 0, Col: 4},
		CellA5:          {Row: 0, Col: 5},
		CellA6:          {Row: 0, Col: 6},
		CellB0:          {Row: 1, Col: 0},
		CellB1:          {Row: 1, Col: 1},
		CellB2:          {Row: 1, Col: 2},
		CellB3:          {Row: 1, Col: 3},
		CellB4:          {Row: 1, Col: 4},
		CellB5:          {Row: 1, Col: 5},
		CellB6:          {Row: 1, Col: 6},
		CellC0:          {Row: 2, Col: 0},
		CellC1:          {Row: 2, Col: 1},
		CellC2:          {Row: 2, Col: 2},
		CellC3:          {Row: 2, Col: 3},
		CellC4:          {Row: 2, Col: 4},
		CellC5:          {Row: 2, Col: 5},
		CellC6:          {Row: 2, Col: 6},
		CellD0:          {Row: 3, Col: 0},
		CellD1:          {Row: 3, Col: 1},
		CellD2:          {Row: 3, Col: 2},
		CellD3:          {Row: 3, Col: 3},
		CellD4:          {Row: 3, Col: 4},
		CellD5:          {Row: 3, Col: 5},
		CellD6:          {Row: 3, Col: 6},
		CellE0:          {Row: 4, Col: 0},
		CellE1:          {Row: 4, Col: 1},
		CellE2:          {Row: 4, Col: 2},
		CellE3:          {Row: 4, Col: 3},
		CellE4:          {Row: 4, Col: 4},
		CellE5:          {Row: 4, Col: 5},
		CellE6:          {Row: 4, Col: 6},
	}
}

func GetCellID(row, col int) CellID {
	for id, pos := range mapCellIDToPosition() {
		if pos.Row == row && pos.Col == col {
			return CellID(id)
		}
	}

	return CellOutOfBounds
}
