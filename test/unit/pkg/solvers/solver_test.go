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
	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)
	assert.Nil(t, err)
	isValidLayout, err := solvers.CheckBoardHasValidLayout(board)
	assert.True(t, isValidLayout)
	assert.Nil(t, err)

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
	brokenBoard, err := boards.NewBoard(cells.NewFactory(), brokenSudoku)

	assert.Nil(t, err)

	invalidLayout, err := solvers.CheckBoardHasValidLayout(brokenBoard)
	assert.False(t, invalidLayout)
	assert.Error(t, err)

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
	brokenBoard2, err := boards.NewBoard(cells.NewFactory(), hardSudoku2)

	assert.Nil(t, err)

	invalidLayout2, err := solvers.CheckBoardHasValidLayout(brokenBoard2)
	assert.False(t, invalidLayout2)
	assert.Error(t, err)
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

	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)

	isValid, err := solvers.CheckValid(board, 0)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 10)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 20)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 30)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 40)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 50)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 60)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 70)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 80)
	assert.True(t, isValid)
	assert.Nil(t, err)
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

	board, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)

	isValid, err := solvers.CheckValid(board, 0)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 10)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 20)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 30)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 40)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 50)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 60)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 70)
	assert.True(t, isValid)
	assert.Nil(t, err)
	isValid, err = solvers.CheckValid(board, 80)
	assert.True(t, isValid)
	assert.Nil(t, err)
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

	easyBoard, err := boards.NewBoard(cells.NewFactory(), easySudoku)

	assert.Nil(t, err)

	solvedEasyBoard, err := solvers.Solve(easyBoard)
	assert.Nil(t, err)

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

	hardBoard, err := boards.NewBoard(cells.NewFactory(), hardSudoku)

	solvedHardBoard, err := solvers.Solve(hardBoard)
	assert.Nil(t, err)

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
