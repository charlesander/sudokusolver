package main

import (
	"github.com/stretchr/testify/assert"
	board "sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"testing"
)

func TestBoardInitiation(t *testing.T) {
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

	board := board.NewBoard(cellFactory, easySudoku)

	cell := board.GetCell(0)
	assert.Equal(t, cell.GetCellValue(), 0)
	assert.Equal(t, cell.GetCellType(), "Settable")

	cell2 := board.GetCell(2)
	assert.Equal(t, cell2.GetCellValue(), 3)
	assert.Equal(t, cell2.GetCellType(), "Preset")

	cell3 := board.GetCell(17)
	assert.Equal(t, cell3.GetCellValue(), 1)
	assert.Equal(t, cell3.GetCellType(), "Preset")
}

func TestSettingCell(t *testing.T) {

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

	board := board.NewBoard(cellFactory, easySudoku)

	//Correctly set a cell value for the settable cell
	board.SetCellValue(0, 1)
	assert.Equal(t, board.GetCell(0).GetCellValue(), 0)
	assert.Equal(t, board.GetCell(0).GetCellType(), "Settable")

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
		cells = append(cells, cellFactory.NewCell(initialValue))
	}

	board := board.NewBoard(cellFactory, easySudoku)

	assert.Equal(t, cellsP, board.GetCells())
}

func TestCheckComplete(t *testing.T) {
	t.Skip("Skipped - implement TestCheckComplete when function implemented")
}