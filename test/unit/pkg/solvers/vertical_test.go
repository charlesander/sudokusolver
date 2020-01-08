package main

import (
	"github.com/stretchr/testify/assert"
	board "sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	"testing"
)

func TestCheckVerticalWithCompletedSudokuNoDuplicates(t *testing.T) {
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

	check, err := solvers.CheckVertical(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 10)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 20)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 30)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 40)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 50)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 60)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 70)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 80)
	assert.True(t, check)
	assert.Nil(t, err)
}

func TestCheckCheckVerticalWithDuplicateOnCol1(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,5,3,2,5,7,1,9,6,//duplicate here
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckVertical(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 1)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 2)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 3)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 4)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 5)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 6)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 7)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 8)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 9)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 10)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 20)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 30)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 40)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 50)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 60)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 70)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 80)
	assert.False(t, check)
	assert.Nil(t, err)

	// Out of allowed range values
	check, err = solvers.CheckVertical(board, 81)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, -1)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, 823421)
	assert.False(t, check)
	assert.Error(t, err)
}

func TestCheckCheckVerticalWithDuplicateOnCol5(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,4,3,6,8,//Duplicate here
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckVertical(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 1)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 2)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 3)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 4)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 5)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 6)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 7)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 8)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 9)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 10)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 20)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 30)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 40)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 50)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 60)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 70)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 80)
	assert.False(t, check)
	assert.Nil(t, err)

	// Out of allowed range values
	check, err = solvers.CheckVertical(board, 81)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, -1)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, 823421)
	assert.False(t, check)
	assert.Error(t, err)
}

func TestCheckCheckVerticalWithDuplicateOnCol7(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		4,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,2,//duplicate here
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckVertical(board, 0)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 1)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 2)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 3)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 4)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 5)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 6)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 7)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 8)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 9)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 10)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 20)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 30)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 40)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 50)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 60)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 70)
	assert.True(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 80)
	assert.False(t, check)
	assert.Nil(t, err)

	// Out of allowed range values
	check, err = solvers.CheckVertical(board, 81)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, -1)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, 823421)
	assert.False(t, check)
	assert.Error(t, err)

}

func TestCheckCheckVerticalWithDuplicateOnCol0(t *testing.T) {
	var complededSudoku = []int{
		9,5,7,6,1,3,2,8,4,
		9,8,3,2,5,7,1,9,6,
		6,1,2,8,4,9,5,3,7,
		1,7,8,3,6,4,9,5,2,
		5,2,4,9,7,1,3,6,8,
		3,6,9,5,2,8,7,4,1,
		8,4,5,7,9,2,6,1,3,
		2,9,1,4,3,6,8,7,5,
		7,3,6,1,8,5,4,2,9}
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	check, err := solvers.CheckVertical(board, 0)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 1)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 2)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 3)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 4)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 5)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 6)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 7)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 8)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 9)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 10)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 20)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 30)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 40)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 50)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 60)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 70)
	assert.False(t, check)
	assert.Nil(t, err)
	check, err = solvers.CheckVertical(board, 80)
	assert.False(t, check)
	assert.Nil(t, err)

	// Out of allowed range values
	check, err = solvers.CheckVertical(board, 81)
	assert.False(t, check)
	assert.Error(t, err)

	check, err = solvers.CheckVertical(board, -1)
	assert.False(t, check)
	assert.Error(t, err)
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
	board, err := board.NewBoard(cells.NewFactory(), complededSudoku)

	assert.Nil(t, err)

	expectedCol := []int{9, 4, 6, 1, 5, 3, 8, 2, 7}
	cols, err := solvers.ExtractVerticalCol(board, 0);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{5, 8, 1, 7, 2, 6, 4, 9, 3}
	cols, err = solvers.ExtractVerticalCol(board, 1);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{7, 3, 2, 8, 4, 9, 5, 1, 6}
	cols, err = solvers.ExtractVerticalCol(board, 2);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{6, 2, 8, 3, 9, 5, 7, 4, 1}
	cols, err = solvers.ExtractVerticalCol(board, 3);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{1, 5, 4, 6, 7, 2, 9, 3, 8}
	cols, err = solvers.ExtractVerticalCol(board, 4);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{3, 7, 9, 4, 1, 8, 2, 6, 5}
	cols, err = solvers.ExtractVerticalCol(board, 5);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{2, 1, 5, 9, 3, 7, 6, 8, 4}
	cols, err = solvers.ExtractVerticalCol(board, 6);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{8, 9, 3, 5, 6, 4, 1, 7, 2}
	cols, err = solvers.ExtractVerticalCol(board, 7);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{4, 6, 7, 2, 8, 1, 3, 5, 9}
	cols, err = solvers.ExtractVerticalCol(board, 8);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{9, 4, 6, 1, 5, 3, 8, 2, 7}
	cols, err = solvers.ExtractVerticalCol(board, 9);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{5, 8, 1, 7, 2, 6, 4, 9, 3}
	cols, err = solvers.ExtractVerticalCol(board, 10);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{7, 3, 2, 8, 4, 9, 5, 1, 6}
	cols, err = solvers.ExtractVerticalCol(board, 11);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{6, 2, 8, 3, 9, 5, 7, 4, 1}
	cols, err = solvers.ExtractVerticalCol(board, 12);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{1, 5, 4, 6, 7, 2, 9, 3, 8}
	cols, err = solvers.ExtractVerticalCol(board, 13);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{3, 7, 9, 4, 1, 8, 2, 6, 5}
	cols, err = solvers.ExtractVerticalCol(board, 14);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{2, 1, 5, 9, 3, 7, 6, 8, 4}
	cols, err = solvers.ExtractVerticalCol(board, 15);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{8, 9, 3, 5, 6, 4, 1, 7, 2}
	cols, err = solvers.ExtractVerticalCol(board, 16);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{4, 6, 7, 2, 8, 1, 3, 5, 9}
	cols, err = solvers.ExtractVerticalCol(board, 17);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{9, 4, 6, 1, 5, 3, 8, 2, 7}
	cols, err = solvers.ExtractVerticalCol(board, 18);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{5, 8, 1, 7, 2, 6, 4, 9, 3}
	cols, err = solvers.ExtractVerticalCol(board, 19);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{7, 3, 2, 8, 4, 9, 5, 1, 6}
	cols, err = solvers.ExtractVerticalCol(board, 20);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{6, 2, 8, 3, 9, 5, 7, 4, 1}
	cols, err = solvers.ExtractVerticalCol(board, 21);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)
	expectedCol = []int{4, 6, 7, 2, 8, 1, 3, 5, 9}
	cols, err = solvers.ExtractVerticalCol(board, 80);
	assert.Nil(t, err)
	assert.Equal(t, expectedCol, cols)

	// Out of allowed range values
	assert.Panics(t, func() {
		solvers.ExtractVerticalCol(board, 81)
	})

	assert.Panics(t, func() {
		solvers.ExtractVerticalCol(board, -1)
	})
}
func TestGetVerticalOffset(t *testing.T) {

	offset, err := solvers.GetVerticalOffset(0)
	assert.Equal(t, 0, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(1)
	assert.Equal(t, 1, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(2)
	assert.Equal(t, 2, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(3)
	assert.Equal(t, 3, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(4)
	assert.Equal(t, 4, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(5)
	assert.Equal(t, 5, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(6)
	assert.Equal(t, 6, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(7)
	assert.Equal(t, 7, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(8)
	assert.Equal(t, 8, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(9)
	assert.Equal(t, 0, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(10)
	assert.Equal(t, 1, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(11)
	assert.Equal(t, 2, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(12)
	assert.Equal(t, 3, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(13)
	assert.Equal(t, 4, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(14)
	assert.Equal(t, 5, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(15)
	assert.Equal(t, 6, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(16)
	assert.Equal(t, 7, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(17)
	assert.Equal(t, 8, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(18)
	assert.Equal(t, 0, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(19)
	assert.Equal(t, 1, offset)
	assert.Nil(t, err)
	offset, err = solvers.GetVerticalOffset(20)
	assert.Equal(t, 2, offset)
	assert.Nil(t, err)
}
