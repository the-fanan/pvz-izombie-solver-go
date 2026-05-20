package simulator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type position struct {
	cellID CellID
	subMatrixRow int
	subMatrixCol int
	subMatrixAltitude int
}

func TestDefaultZombieMovement(t *testing.T) {
	tests := []struct {
		resolution int
		rows int
		cols int
		startingCellID CellID
		startingSubmatrixRow int
		startingSubmatrixCol int
		expected []position
	}{
		{
			resolution: 3,
			rows: 1,
			cols: 2,
			startingCellID: CellA1,
			startingSubmatrixRow: 1,
			startingSubmatrixCol: 1,
			expected: []position{
				{
					cellID: CellA1,
					subMatrixRow: 1,
					subMatrixCol: 0,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 1,
					subMatrixCol: 2,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 1,
					subMatrixCol: 1,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 1,
					subMatrixCol: 0,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellOutOfBounds,
					subMatrixRow: 1,
					subMatrixCol: -1,
					subMatrixAltitude: 0,
				},
			},
		},
		{
			resolution: 9,
			rows: 1,
			cols: 1,
			startingCellID: CellA0,
			startingSubmatrixRow: 4,
			startingSubmatrixCol: 4,
			expected: []position{
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 4,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 4,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 3,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 3,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 3,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 2,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 2,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 2,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 1,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 1,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 1,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 0,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 0,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellA0,
					subMatrixRow: 4,
					subMatrixCol: 0,
					subMatrixAltitude: 0,
				},
				{
					cellID: CellOutOfBounds,
					subMatrixRow: 4,
					subMatrixCol: -1,
					subMatrixAltitude: 0,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("move works for normal zombie in resolution %d for %d rows and %d cols", test.resolution, test.rows, test.cols), func(t *testing.T){
			frame := NewFrame(test.rows, test.cols, test.resolution, 0, 100)
			zombie := NewZombie(ZombieImp, test.resolution, test.startingCellID, test.startingSubmatrixRow, test.startingSubmatrixCol, 0)

			frame.SetCellZombie(zombie, test.startingCellID.Position().Row, test.startingCellID.Position().Col)

			for _, p := range test.expected {
				zombie.Move(frame)
				c, sr, sc, sa := zombie.GetCurrentPosition()
				assert.Equal(t, p.cellID, c, fmt.Sprintf("cell is mismatched. expected %s but got %s", p.cellID.String(), c.String()))
				assert.Equal(t, p.subMatrixRow, sr, fmt.Sprintf("submatrix row is mismatched. expected %d but got %d", p.subMatrixRow, sr))
				assert.Equal(t, p.subMatrixCol, sc,  fmt.Sprintf("submatrix column is mismatched. expected %d but got %d", p.subMatrixCol, sc))
				assert.Equal(t, p.subMatrixAltitude, sa, fmt.Sprintf("submatrix altitude is mismatched. expected %d but got %d", p.subMatrixAltitude, sa))
			}
		})
	}
}
