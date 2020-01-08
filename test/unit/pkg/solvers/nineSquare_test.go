package main

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckNineSquaresAllCorrectCells(t *testing.T) {
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
	board, err := boards.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckNineSquares(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)

	check, err = solvers.CheckNineSquares(board, 1)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 10)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 20)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 30)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 40)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 50)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 60)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 70)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 80)
	assert.True(t, check)
	assert.Nil(t, err)

	// Out of expected range error
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, 81)
	})
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, -1)
	})
}

func TestCheckNineSquaresDuplicateBottomRight(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,2,5,
		7,3,6,1,8,5,4,2,9}
	board, err := boards.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckNineSquares(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)

	check, err = solvers.CheckNineSquares(board, 1)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 10)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 20)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 30)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 40)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 50)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 60)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 70)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 80)
	assert.True(t, check)
	assert.Nil(t, err)

	// Out of expected range error
	assert.Panics(t, func() {
		solvers.CheckNineSquares(board, 81)
	})
	assert.Panics(t, func() {
		solvers.CheckNineSquares(board, -1)
	})
}

func TestCheckNineSquaresDuplicateTopRight(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,8,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := boards.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckNineSquares(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)

	check, err = solvers.CheckNineSquares(board, 1)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 10)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 20)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 30)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 40)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 50)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 60)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 70)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckNineSquares(board, 80)
	assert.True(t, check)
	assert.Nil(t, err)

	// Out of expected range error
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, 81)
	})
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, -1)
	})
}

func TestGetSudokuSquareIndexes(t *testing.T) {
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
	board, err := boards.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	expectedSquareIndexes := []int{0,1,2,9,10,11,18,19,20}

	squareIndexes, err := solvers.GetSudokuSquareIndexes(board, 0);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 1);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 2);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 9);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 10);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 11);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 18);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 19);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 20);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)

	// Top center
	expectedSquareIndexes = []int{3,4,5,12,13,14,21,22,23}

	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 3);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 4);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 5);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 12);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 13);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 14);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 21);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 22);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 23);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)

	// Top right
	expectedSquareIndexes = []int{6,7,8,15,16,17,24,25,26}

	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 6);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 7);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 8);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 15);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 16);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 17);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 24);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 25);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 26);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)

	// Bottom right
	expectedSquareIndexes = []int{60,61,62,69,70,71,78,79,80}

	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 60);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 61);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 62);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 69);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 70);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 71);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 78);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 79);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	squareIndexes, err = solvers.GetSudokuSquareIndexes(board, 80);
	assert.Equal(t, expectedSquareIndexes, squareIndexes)
	assert.Nil(t, err)
	// Out of expected range error
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, 81)
	})
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, -1)
	})
}

func TestGetSudokuSquareValues(t *testing.T) {
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
	board, err := boards.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	// Top left
	squareIndexes := []int{0,1,2,9,10,11,18,19,20}
	expectedSquareValues := []int{9,5,7,4,8,3,6,1,2}
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))

	// Top center
	squareIndexes = []int{3,4,5,12,13,14,21,22,23}
	expectedSquareValues = []int{6,1,3,2,5,7,8,4,9}

	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))

	// Top right
	squareIndexes = []int{6,7,8,15,16,17,24,25,26}
	expectedSquareValues = []int{2,8,4,1,9,6,5,3,7}

	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))

	// Bottom right
	squareIndexes = []int{60,61,62,69,70,71,78,79,80}
	expectedSquareValues = []int{6,1,3,8,7,5,4,2,9}

	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))
	assert.Equal(t, expectedSquareValues, solvers.GetSudokuSquareValues(board, squareIndexes))

	squareIndexes = []int{6,1,83}
	// Out of expected range error
	assert.Panics(t, func() {
		solvers.GetSudokuSquareValues(board, squareIndexes)
	})

	squareIndexes = []int{6,1,-1,234}
	assert.Panics(t, func() {
		solvers.GetSudokuSquareValues(board, squareIndexes)
	})
}

func TestExtractNineSquares(t *testing.T) {

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
	board, err := boards.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	// Top left
	expectedSquare := []int{9,5,7,4,8,3,6,1,2}

	squares, err := solvers.ExtractNineSquares(board, 0)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 1)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 2)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 9)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 10)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 11)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 18)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 19)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 20)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)

	// Top center
	expectedSquare = []int{6,1,3,2,5,7,8,4,9}

	squares, err = solvers.ExtractNineSquares(board, 3)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 4)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 5)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 12)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 13)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 14)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 21)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 22)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 23)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)

	// Top right
	expectedSquare = []int{2,8,4,1,9,6,5,3,7}

	squares, err = solvers.ExtractNineSquares(board, 6)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 7)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 8)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 15)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 16)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 17)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 24)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 25)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 26)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)

	// Bottom right
	expectedSquare = []int{6,1,3,8,7,5,4,2,9}

	squares, err = solvers.ExtractNineSquares(board, 60)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 61)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 62)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 69)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 70)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 71)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 78)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 79)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)
	squares, err = solvers.ExtractNineSquares(board, 80)
	assert.Equal(t, expectedSquare, squares)
	assert.Nil(t, err)

	// Out of expected range error
	squares, err = solvers.ExtractNineSquares(board, 80)
	assert.Nil(t, squares)
	assert.Error(t, err)

	squares, err = solvers.ExtractNineSquares(board, -1)
	assert.Nil(t, squares)
	assert.Error(t, err)
}