package main

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckNineSquares(t *testing.T) {
	t.Skip("Skipped - implement TestCheckNineSquares when function implemented")
}

func TestExtractNineSquares(t *testing.T) {
	t.Skip("Skipped - implement TestExtractNineSquares when function implemented")
}

func TestGetSquareValues(t *testing.T) {

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
	board := boards.NewBoard(cells.NewFactory(), complededSudoku)

	// Top left
	expectedSquare := []int{9,5,7,4,8,3,6,1,2}
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 0))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 1))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 2))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 9))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 10))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 11))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 18))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 19))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 20))

	// Top center
	expectedSquare = []int{6,1,3,2,5,7,8,4,9}

	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 3))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 4))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 5))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 12))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 13))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 14))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 21))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 22))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 23))

	// Top right
	expectedSquare = []int{2,8,4,1,9,6,5,3,7}

	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 6))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 7))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 8))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 15))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 16))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 17))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 24))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 25))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 26))

	// Bottom right
	expectedSquare = []int{6,1,3,8,7,5,4,2,9}

	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 60))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 61))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 62))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 69))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 70))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 71))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 78))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 79))
	assert.Equal(t, expectedSquare, solvers.ExtractNineSquares(board, 80))

	// Out of expected range error
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, 81)
	})
	assert.Panics(t, func() {
		solvers.ExtractNineSquares(board, -1)
	})
}