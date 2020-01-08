package main

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"testing"
)

func TestBoardInitiation(t *testing.T) {

	var easySudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 0, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}

	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)

	cell, err := board.GetCell(0)
	assert.Equal(t, cell.GetCellValue(), 0)
	assert.Equal(t, cell.GetCellType(), "Settable")
	assert.Nil(t, err)

	cell2, err := board.GetCell(2)
	assert.Equal(t, cell2.GetCellValue(), 3)
	assert.Equal(t, cell2.GetCellType(), "Preset")
	assert.Nil(t, err)

	cell3, err := board.GetCell(17)
	assert.Equal(t, cell3.GetCellValue(), 1)
	assert.Equal(t, cell3.GetCellType(), "Preset")
	assert.Nil(t, err)
}

func TestSettingCell(t *testing.T) {
	t.Skip("Test when proper error handling implemented")
	var easySudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 0, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}

	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)

	//Correctly set a cell value for the settable cell
	board.SetCellValue(0, 1)
	cell, err := board.GetCell(0)
	assert.Nil(t, err)
	assert.Equal(t, cell.GetCellValue(), 1)
	assert.Equal(t, cell.GetCellType(), "Settable")

	board.SetCellValue(0, 2)
	cell, err = board.GetCell(0)
	assert.Nil(t, err)
	assert.Equal(t, cell.GetCellValue(), 2)

	board.SetCellValue(0, 3)
	cell, err = board.GetCell(0)
	assert.Nil(t, err)
	assert.Equal(t, cell.GetCellValue(), 3)

	board.SetCellValue(0, 5)
	cell, err = board.GetCell(0)
	assert.Nil(t, err)
	assert.Equal(t, cell.GetCellValue(), 5)

	board.SetCellValue(0, boards.BOARD_SIDE_LENGTH)
	cell, err = board.GetCell(0)
	assert.Nil(t, err)
	assert.Equal(t, cell.GetCellValue(), 9)
	//Fail setting a cell value for a settable cell

	assert.Panics(t, func() {
		board.SetCellValue(0, -1)
	})
	assert.Panics(t, func() {
		board.SetCellValue(0, 10)
	})

	//Fail setting a cell value for a preset cell
	assert.Panics(t, func() {
		board.SetCellValue(36, 1)
	})


	assert.Panics(t, func() {
		board.SetCellValue(78, -1)
	})
	assert.Panics(t, func() {
		board.SetCellValue(78, 10)
	})
}

func TestGetCells(t *testing.T) {

	cellFactory := cells.NewFactory()

	var easySudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 0, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}

	var cells []cells.IndividualCell
	cellsP := &cells
	for _, initialValue := range easySudoku {
		newCell, _ := cellFactory.NewCell(initialValue)
		cells = append(cells, newCell)
	}

	board, err := boards.NewBoard(cellFactory, easySudoku)

	assert.Nil(t, err)

	assert.Equal(t, cellsP, board.GetCells())
}

func TestSetComplete(t *testing.T) {
	var easySudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 0, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}

	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)
	assert.False(t, board.CheckComplete())
	board.SetComplete()
	assert.True(t, board.CheckComplete())
}

func TestCheckComplete(t *testing.T) {
	var easySudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 0, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}

	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)
	assert.False(t, board.CheckComplete())
	board.SetComplete()
	assert.True(t, board.CheckComplete())
}
