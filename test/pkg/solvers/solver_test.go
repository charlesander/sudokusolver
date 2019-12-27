package main

import (
	"github.com/stretchr/testify/assert"
	board "sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckValid(t *testing.T) {
	t.Skip("Skipped - implement TestCheckValid when function implemented")
}

func TestCheckHorizonal(t *testing.T) {
	t.Skip("Skipped - implement TestCheckHorizonal when function implemented")
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
	board := board.NewBoard(cells.NewFactory(), easySudoku)

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

func TestCheckVertical(t *testing.T) {
	t.Skip("Skipped - implement TestCheckVertical when function implemented")
}

func TestCheckSquares(t *testing.T) {
	t.Skip("Skipped - implement TestCheckSquares when function implemented")
}

func TestSolve(t *testing.T) {
	t.Skip("Skipped - implement TestSolve when function implemented")
}