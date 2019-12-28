package main

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckBoardHasValidLayout(t *testing.T) {
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
	board := boards.NewBoard(cells.NewFactory(), easySudoku)
	assert.True(t, solvers.CheckBoardHasValidLayout(board))

	var brokenSudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 2, 2, 9, 0, 0,
		7, 0, 0, 0, 2, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}
	brokenBoard := boards.NewBoard(cells.NewFactory(), brokenSudoku)
	assert.False(t, solvers.CheckBoardHasValidLayout(brokenBoard))

	var hardSudoku2 = []int{
		6, 0, 0, 0, 0, 0, 1, 5, 0,
		9, 5, 4, 7, 1, 0, 0, 8, 0,
		0, 0, 0, 5, 0, 2, 6, 0, 0,

		8, 0, 0, 0, 9, 4, 0, 0, 6,
		0, 0, 3, 8, 0, 5, 4, 0, 0,
		4, 0, 0, 3, 7, 0, 0, 0, 8,

		0, 0, 6, 9, 0, 3, 0, 0, 0,
		0, 2, 0, 0, 4, 7, 8, 9, 3,
		0, 4, 9, 0, 0, 0, 0, 9, 5,
	}
	brokenBoard2 := boards.NewBoard(cells.NewFactory(), hardSudoku2)
	assert.False(t, solvers.CheckBoardHasValidLayout(brokenBoard2))
}
func TestCheckValidWithDuplicate(t *testing.T) {
	var easySudoku = []int{
		0, 0, 3, 0, 2, 0, 6, 0, 0,
		9, 0, 0, 3, 0, 5, 0, 0, 1,
		0, 0, 1, 8, 0, 6, 4, 0, 0,

		0, 0, 8, 1, 2, 2, 9, 0, 0,
		7, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 7, 0, 8, 2, 0, 0,

		0, 0, 2, 6, 0, 9, 5, 0, 0,
		8, 0, 0, 2, 0, 3, 0, 0, 0,
		0, 0, 5, 0, 1, 0, 3, 0, 0,
	}

	board := boards.NewBoard(cells.NewFactory(), easySudoku)
	assert.True(t, solvers.CheckValid(board, 0))
	assert.True(t, solvers.CheckValid(board, 10))
	assert.True(t, solvers.CheckValid(board, 20))
	assert.False(t, solvers.CheckValid(board, 30))
	assert.False(t, solvers.CheckValid(board, 40))
	assert.False(t, solvers.CheckValid(board, 50))
	assert.False(t, solvers.CheckValid(board, 60))
	assert.False(t, solvers.CheckValid(board, 70))
	assert.False(t, solvers.CheckValid(board, 80))
}

func TestCheckValid(t *testing.T) {
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

	board := boards.NewBoard(cells.NewFactory(), easySudoku)
	assert.True(t, solvers.CheckValid(board, 0))
	assert.True(t, solvers.CheckValid(board, 10))
	assert.True(t, solvers.CheckValid(board, 20))
	assert.True(t, solvers.CheckValid(board, 30))
	assert.True(t, solvers.CheckValid(board, 40))
	assert.True(t, solvers.CheckValid(board, 50))
	assert.True(t, solvers.CheckValid(board, 60))
	assert.True(t, solvers.CheckValid(board, 70))
	assert.True(t, solvers.CheckValid(board, 80))
}

func TestSolve(t *testing.T) {
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

	easyBoard := boards.NewBoard(cells.NewFactory(), easySudoku)

	solvedEasyBoard := solvers.Solve(easyBoard)

	var solvedEasyBoardExpectedResult = []int{
		4, 8, 3, 9, 2, 1, 6, 5, 7,
		9, 6, 7, 3, 4, 5, 8, 2, 1,
		2, 5, 1, 8, 7, 6, 4, 3, 9,
		5, 4, 8, 1, 6, 2, 9, 7, 3,
		7, 2, 9, 5, 3, 4, 1, 6, 8,
		1, 3, 6, 7, 9, 8, 2, 4, 5,
		3, 7, 2, 6, 8, 9, 5, 1, 4,
		8, 1, 4, 2, 5, 3, 7, 9, 6,
		6, 9, 5, 4, 1, 7, 3, 8, 2}

	for i, element := range solvedEasyBoardExpectedResult {
		assert.Equal(t, element, solvedEasyBoard.GetCell(i).GetCellValue())
	}

	var hardSudoku = []int{
		6, 0, 0, 0, 0, 0, 1, 5, 0,
		9, 5, 4, 7, 1, 0, 0, 8, 0,
		0, 0, 0, 5, 0, 2, 6, 0, 0,

		8, 0, 0, 0, 9, 4, 0, 0, 6,
		0, 0, 3, 8, 0, 5, 4, 0, 0,
		4, 0, 0, 3, 7, 0, 0, 0, 8,

		0, 0, 6, 9, 0, 3, 0, 0, 0,
		0, 2, 0, 0, 4, 7, 8, 9, 3,
		0, 4, 9, 0, 0, 0, 0, 0, 5,
	}

	hardBoard := boards.NewBoard(cells.NewFactory(), hardSudoku)

	solvedHardBoard := solvers.Solve(hardBoard)

	var solvedHardBoardExpectedResult = []int{
		6, 3, 2, 4, 8, 9, 1, 5, 7,
		9, 5, 4, 7, 1, 6, 3, 8, 2,
		1, 7, 8, 5, 3, 2, 6, 4, 9,
		8, 1, 7, 2, 9, 4, 5, 3, 6,
		2, 9, 3, 8, 6, 5, 4, 7, 1,
		4, 6, 5, 3, 7, 1, 9, 2, 8,
		7, 8, 6, 9, 5, 3, 2, 1, 4,
		5, 2, 1, 6, 4, 7, 8, 9, 3,
		3, 4, 9, 1, 2, 8, 7, 6, 5}

	for i, element := range solvedHardBoardExpectedResult {
		assert.Equal(t, element, solvedHardBoard.GetCell(i).GetCellValue())
	}

}
