package solvers

import (
	"fmt"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	utilties "sudokusolver/pkg/utilites"
)

func checkValid(board boards.Board, index int) bool {
	//check horizontal
	//check vertical
	//check 3x3 squares
	return true
}

func CheckHorizontal(board boards.Board, index int) bool {
	for i := 0; i < boards.CELL_COUNT; i += 9 {
		if i <= index {
			if !utilties.AreSudokuValuesUnique(ExtractHorizontalRow(board, i)) {
				return false
			}
		} else {
			break
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
	return index % 9
}

func CheckVertical(board boards.Board, index int) bool {
	//col := []int{}
	//currentOffset := GetVerticalOffset(index)
	//return col
	return true
}

func ExtractVerticalCol(board boards.Board, offset int) []int {
	col := []int{}
	i := offset
	for i <= boards.CELL_COUNT {
		col = append(col, board.GetCell(i).GetCellValue())
		i+= 9
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
		checkValid(board, i)
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
	}

	return board
}
