package main

import (
	"github.com/stretchr/testify/assert"
	board "sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckVerticalWithCompletedSudokuNoDuplicates(t *testing.T) {
	t.Skip("Skipped - implement TestCheckVerticalWithCompletedSudokuNoDuplicates when function implemented")
}

func TestExtractVerticalCol(t *testing.T) {
	var complededSudoku = []int{
		9, 5, 7, 6, 1, 3, 2, 8, 4,
		4, 8, 3, 2, 5, 7, 1, 9, 6,
		6, 1, 2, 8, 4, 9, 5, 3, 7,
		1, 7, 8, 3, 6, 4, 9, 5, 2,
		5, 2, 4, 9, 7, 1, 3, 6, 8,
		3, 6, 9, 5, 2, 8, 7, 4, 1,
		8, 4, 5, 7, 9, 2, 6, 1, 3,
		2, 9, 1, 4, 3, 6, 8, 7, 5,
		7, 3, 6, 1, 8, 5, 4, 2, 9}
	board := board.NewBoard(cells.NewFactory(), complededSudoku)

	expectedCol := []int{9, 4, 6, 1, 5, 3, 8, 2, 7}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 0))
	expectedCol = []int{5, 8, 1, 7, 2, 6, 4, 9, 3}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 1))
	expectedCol = []int{7, 3, 2, 8, 4, 9, 5, 1, 6}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 2))
	expectedCol = []int{6, 2, 8, 3, 9, 5, 7, 4, 1}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 3))
	expectedCol = []int{1, 5, 4, 6, 7, 2, 9, 3, 8}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 4))
	expectedCol = []int{3, 7, 9, 4, 1, 8, 2, 6, 5}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 5))
	expectedCol = []int{2, 1, 5, 9, 3, 7, 6, 8, 4}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 6))
	expectedCol = []int{8, 9, 3, 5, 6, 4, 1, 7, 2}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 7))
	expectedCol = []int{4, 6, 7, 2, 8, 1, 3, 5, 9}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 8))
	expectedCol = []int{9, 4, 6, 1, 5, 3, 8, 2, 7}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 9))
	expectedCol = []int{5, 8, 1, 7, 2, 6, 4, 9, 3}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 10))
	expectedCol = []int{7, 3, 2, 8, 4, 9, 5, 1, 6}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 11))
	expectedCol = []int{6, 2, 8, 3, 9, 5, 7, 4, 1}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 12))
	expectedCol = []int{1, 5, 4, 6, 7, 2, 9, 3, 8}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 13))
	expectedCol = []int{3, 7, 9, 4, 1, 8, 2, 6, 5}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 14))
	expectedCol = []int{2, 1, 5, 9, 3, 7, 6, 8, 4}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 15))
	expectedCol = []int{8, 9, 3, 5, 6, 4, 1, 7, 2}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 16))
	expectedCol = []int{4, 6, 7, 2, 8, 1, 3, 5, 9}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 17))
	expectedCol = []int{9, 4, 6, 1, 5, 3, 8, 2, 7}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 18))
	expectedCol = []int{5, 8, 1, 7, 2, 6, 4, 9, 3}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 19))
	expectedCol = []int{7, 3, 2, 8, 4, 9, 5, 1, 6}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 20))
	expectedCol = []int{6, 2, 8, 3, 9, 5, 7, 4, 1}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 21))
	expectedCol = []int{4, 6, 7, 2, 8, 1, 3, 5, 9}
	assert.Equal(t, expectedCol, solvers.ExtractVerticalCol(board, 80))

	// Out of allowed range values
	assert.Panics(t, func() {
		solvers.ExtractVerticalCol(board, 81)
	})

	assert.Panics(t, func() {
		solvers.ExtractVerticalCol(board, -1)
	})
}
func TestGetVerticalOffset(t *testing.T) {

	assert.Equal(t, 0, solvers.GetVerticalOffset(0))
	assert.Equal(t, 1, solvers.GetVerticalOffset(1))
	assert.Equal(t, 2, solvers.GetVerticalOffset(2))
	assert.Equal(t, 3, solvers.GetVerticalOffset(3))
	assert.Equal(t, 4, solvers.GetVerticalOffset(4))
	assert.Equal(t, 5, solvers.GetVerticalOffset(5))
	assert.Equal(t, 6, solvers.GetVerticalOffset(6))
	assert.Equal(t, 7, solvers.GetVerticalOffset(7))
	assert.Equal(t, 8, solvers.GetVerticalOffset(8))
	assert.Equal(t, 0, solvers.GetVerticalOffset(9))
	assert.Equal(t, 1, solvers.GetVerticalOffset(10))
	assert.Equal(t, 2, solvers.GetVerticalOffset(11))
	assert.Equal(t, 3, solvers.GetVerticalOffset(12))
	assert.Equal(t, 4, solvers.GetVerticalOffset(13))
	assert.Equal(t, 5, solvers.GetVerticalOffset(14))
	assert.Equal(t, 6, solvers.GetVerticalOffset(15))
	assert.Equal(t, 7, solvers.GetVerticalOffset(16))
	assert.Equal(t, 8, solvers.GetVerticalOffset(17))
	assert.Equal(t, 0, solvers.GetVerticalOffset(18))
	assert.Equal(t, 1, solvers.GetVerticalOffset(19))
	assert.Equal(t, 2, solvers.GetVerticalOffset(20))
}
