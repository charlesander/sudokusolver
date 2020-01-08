package main

import (
	"github.com/stretchr/testify/assert"
	board "sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckHorizonalWithCompletedSudokuNoDuplicates(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	assert.True(t, solvers.CheckHorizontal(board, 0))
	assert.True(t, solvers.CheckHorizontal(board, 1))
	assert.True(t, solvers.CheckHorizontal(board, 10))
	assert.True(t, solvers.CheckHorizontal(board, 20))
	assert.True(t, solvers.CheckHorizontal(board, 30))
	assert.True(t, solvers.CheckHorizontal(board, 40))
	assert.True(t, solvers.CheckHorizontal(board, 50))
	assert.True(t, solvers.CheckHorizontal(board, 60))
	assert.True(t, solvers.CheckHorizontal(board, 70))
	assert.True(t, solvers.CheckHorizontal(board, 80))
}

func TestCheckHorizonalWithDuplicateOnRow1(t *testing.T) {
	var complededSudoku = []int{
		//9,5,7,6,1,3,2,8,4,
		9,5,7,6,1,3,2,8,9,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	assert.False(t, solvers.CheckHorizontal(board, 0))
	assert.False(t, solvers.CheckHorizontal(board, 1))
	assert.False(t, solvers.CheckHorizontal(board, 2))
	assert.False(t, solvers.CheckHorizontal(board, 3))
	assert.False(t, solvers.CheckHorizontal(board, 1))
	assert.False(t, solvers.CheckHorizontal(board, 6))
	assert.False(t, solvers.CheckHorizontal(board, 7))
	assert.False(t, solvers.CheckHorizontal(board, 8))
	assert.False(t, solvers.CheckHorizontal(board, 9))
	assert.False(t, solvers.CheckHorizontal(board, 10))
	assert.False(t, solvers.CheckHorizontal(board, 20))
	assert.False(t, solvers.CheckHorizontal(board, 30))
	assert.False(t, solvers.CheckHorizontal(board, 40))
	assert.False(t, solvers.CheckHorizontal(board, 50))
	assert.False(t, solvers.CheckHorizontal(board, 60))
	assert.False(t, solvers.CheckHorizontal(board, 70))
	assert.False(t, solvers.CheckHorizontal(board, 80))
	assert.False(t, solvers.CheckHorizontal(board, 81))

}

func TestCheckHorizonalWithDuplicateOnRow5(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		//5,2,4,9,7,1,3,6,8,
		5,5,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	assert.True(t, solvers.CheckHorizontal(board, 0))
	assert.True(t, solvers.CheckHorizontal(board, 1))
	assert.True(t, solvers.CheckHorizontal(board, 10))
	assert.True(t, solvers.CheckHorizontal(board, 20))
	assert.True(t, solvers.CheckHorizontal(board, 35))
	// we're now reaching the erronious row, so it should now fail
	assert.False(t, solvers.CheckHorizontal(board, 36))
	assert.False(t, solvers.CheckHorizontal(board, 40))
	assert.False(t, solvers.CheckHorizontal(board, 50))
	assert.False(t, solvers.CheckHorizontal(board, 65))
	assert.False(t, solvers.CheckHorizontal(board, 70))
	assert.False(t, solvers.CheckHorizontal(board, 80))
}

func TestCheckHorizonalWithDuplicateOnRow9(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		//7,3,6,1,8,5,4,2,9
		7,3,6,1,8,5,4,2,7}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	assert.True(t, solvers.CheckHorizontal(board, 0))
	assert.True(t, solvers.CheckHorizontal(board, 1))
	assert.True(t, solvers.CheckHorizontal(board, 10))
	assert.True(t, solvers.CheckHorizontal(board, 20))
	assert.True(t, solvers.CheckHorizontal(board, 35))

	assert.True(t, solvers.CheckHorizontal(board, 36))
	assert.True(t, solvers.CheckHorizontal(board, 40))
	assert.True(t, solvers.CheckHorizontal(board, 50))
	assert.True(t, solvers.CheckHorizontal(board, 65))
	assert.True(t, solvers.CheckHorizontal(board, 70))
	// we're now reaching the erronious row, so it should now fail
	assert.False(t, solvers.CheckHorizontal(board, 72))
	assert.False(t, solvers.CheckHorizontal(board, 77))
	assert.False(t, solvers.CheckHorizontal(board, 80))
}

func TestExtractHorizontalRow(t *testing.T) {
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
	board, err := board.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)

	// extract 1st row

	expectedRow := []int{0, 0, 3, 0, 2, 0, 6, 0, 0}

	extractedRow := solvers.ExtractHorizontalRow(board, 0)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 1)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 2)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 3)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 4)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 5)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 6)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 7)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 8)
	assert.Equal(t, expectedRow, extractedRow)

	// extract 2nd row

	expectedRow = []int{9, 0, 0, 3, 0, 5, 0, 0, 1}

	extractedRow = solvers.ExtractHorizontalRow(board, 9)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 10)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 17)
	assert.Equal(t, expectedRow, extractedRow)

	// extract 6th row

	expectedRow = []int{0, 0, 6, 7, 0, 8, 2, 0, 0}

	extractedRow = solvers.ExtractHorizontalRow(board, 45)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 46)
	assert.Equal(t, expectedRow, extractedRow)

	// extract 9th row

	expectedRow = []int{0, 0, 5, 0, 1, 0, 3, 0, 0}

	extractedRow = solvers.ExtractHorizontalRow(board, 72)
	assert.Equal(t, expectedRow, extractedRow)

	extractedRow = solvers.ExtractHorizontalRow(board, 80)
	assert.Equal(t, expectedRow, extractedRow)
}
