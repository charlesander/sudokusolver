package solvers

import (
	"fmt"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	utilties "sudokusolver/pkg/utilites"
)

func checkValid(board boards.Board, index int) bool {
	return CheckHorizontal(board, index) &&
		CheckVertical(board, index) &&
		CheckNineSquares(board, index)

}

func CheckNineSquares(board boards.Board, index int) bool {
	return utilties.AreSudokuValuesUnique(ExtractNineSquares(board, index))
}

/**
 * The sudoku is split into nine squares, each three by three.
 * If the idex of the cell being checked falls in one of the squares, return
 * all the values of the cells within the square
 */
func ExtractNineSquares(board boards.Board, index int) []int {
	squareIndexes := GetSudokuSquareIndexes(board, index)
	return GetSudokuSquareValues(board, squareIndexes)
}

func GetSudokuSquareIndexes(board boards.Board, index int) []int {
	// These are the indexes of the cells within each square
	topLeftSquare := []int{0, 1, 2, 9, 10, 11, 18, 19, 20}
	topCentreSquare := []int{3, 4, 5, 12, 13, 14, 21, 22, 23}
	topRightSquare := []int{6, 7, 8, 15, 16, 17, 24, 25, 26}
	middleLeftSquare := []int{27, 28, 29, 36, 37, 38, 45, 46, 47}
	middleCentreSquare := []int{30, 31, 32, 39, 40, 41, 48, 49, 50}
	middleRightSquare := []int{33, 34, 35, 42, 43, 44, 51, 52, 53}
	bottomLeftSquare := []int{54, 55, 56, 63, 64, 65, 72, 73, 74}
	bottomCentreSquare := []int{57, 58, 59, 66, 67, 68, 75, 76, 77}
	bottomRightSquare := []int{60, 61, 62, 69, 70, 71, 78, 79, 80}

	if utilties.IntInSlice(index, topLeftSquare) {
		return topLeftSquare
	} else if utilties.IntInSlice(index, topCentreSquare) {
		return topCentreSquare
	} else if utilties.IntInSlice(index, topRightSquare) {
		return topRightSquare
	} else if utilties.IntInSlice(index, middleLeftSquare) {
		return middleLeftSquare
	} else if utilties.IntInSlice(index, middleCentreSquare) {
		return middleCentreSquare
	} else if utilties.IntInSlice(index, middleRightSquare) {
		return middleRightSquare
	} else if utilties.IntInSlice(index, bottomLeftSquare) {
		return bottomLeftSquare
	} else if utilties.IntInSlice(index, bottomCentreSquare) {
		return bottomCentreSquare
	} else if utilties.IntInSlice(index, bottomRightSquare) {
		return bottomRightSquare
	} else {
		panic("Incorrect index supplied to GetSudokuSquareIndexes")
	}
}

func GetSudokuSquareValues(board boards.Board, squareIndex []int) []int {
	squareValues := []int{}
	for _, i := range squareIndex {
		squareValues = append(squareValues, board.GetCell(i).GetCellValue())
	}
	return squareValues
}

func CheckHorizontal(board boards.Board, index int) bool {
	for i := 0; i < boards.CELL_COUNT; i += 9 {
		if i <= index {
			if !utilties.AreSudokuValuesUnique(ExtractHorizontalRow(board, i)) {
				return false
			}
		} else {
			continue
		}
	}
	return true
}

func ExtractHorizontalRow(board boards.Board, index int) []int {
	row := []int{}
	start := 0
	end := 0
	if index <= 8 {
		start = 0
		end = 8
	} else if index <= 17 {
		start = 9
		end = 17
	} else if index <= 26 {
		start = 18
		end = 26
	} else if index <= 35 {
		start = 27
		end = 35
	} else if index <= 44 {
		start = 36
		end = 44
	} else if index <= 53 {
		start = 45
		end = 53
	} else if index <= 62 {
		start = 54
		end = 62
	} else if index <= 71 {
		start = 63
		end = 71
	} else if index <= 80 {
		start = 72
		end = 80
	} else {
		panic("Incorrect index supplied to ExtractHorizonalRow")
	}

	for i := start; i <= end; i++ {
		row = append(row, board.GetCell(i).GetCellValue())
	}
	return row
}

func GetVerticalOffset(index int) int {
	if index >= boards.CELL_COUNT || index < 0 {
		panic("Incorrect index supplied to GetVerticalOffset")
	}
	return index % 9
}

func CheckVertical(board boards.Board, index int) bool {
	currentOffset := GetVerticalOffset(index)
	for i := 0; i <= currentOffset; i++ {
		if i <= index {
			if !utilties.AreSudokuValuesUnique(ExtractVerticalCol(board, i)) {
				return false
			}
		} else {
			continue
		}
	}
	return true
}

func ExtractVerticalCol(board boards.Board, offset int) []int {
	if offset >= boards.CELL_COUNT || offset < 0 {
		panic("Incorrect index supplied to ExtractHorizonalRow")
	}
	col := []int{}
	i := 0
	i = offset % 9
	for i < boards.CELL_COUNT {
		col = append(col, board.GetCell(i).GetCellValue())
		i += 9
	}
	return col
}

func checkSquares(board boards.Board, index int) bool {
	return true
}

func Solve(board boards.Board) boards.Board {
	fmt.Println(board.GetCell(0).GetCellType())
	i := 0
	for i < boards.CELL_COUNT {
		fmt.Print(i, ",")
		if board.GetCell(i).GetCellType() == cells.PRESET_CELL_TYPE {
			//skip over preset cells (they're not setable)
			i++
			continue
		} else if board.GetCell(i).GetCellValue() == cells.MAX_CELL_VALUE {
			//We've tried all the valid values for this cell, time to backtrack
			board.GetCell(i).SetCellValue(0)
			for {
				i--
				if board.GetCell(i).GetCellType() == cells.SETTABLE_CELL_TYPE {
					continue
				}
			}
		} else {

			if board.GetCell(i).GetCellValue() == 0 {
				board.GetCell(i).SetCellValue(1)
			} else if checkValid(board, i) {
				i++
			} else {
				board.GetCell(i).SetCellValue(board.GetCell(i).GetCellValue() + 1)
			}
		}
	}

	return board
}
	/*
		if (checkValid(board, i)) {
			i++
			continue
		}
		if board.GetCell(i).GetCellType() == cells.PRESET_CELL_TYPE {
			//skip over preset cells (they're not setable)
			i++
			continue
		} else if board.GetCell(i).GetCellValue() == cells.MAX_CELL_VALUE {
			//We've tried all the valid values for this cell, time to backtrack
			board.GetCell(i).SetCellValue(0)
			for {
				i--
				if board.GetCell(i).GetCellType() == cells.SETTABLE_CELL_TYPE {
					continue
				}
			}
		} else {
			board.GetCell(i).SetCellValue(board.GetCell(i).GetCellValue() + 1)
			i++
		}
		//Otherwise
	*/
