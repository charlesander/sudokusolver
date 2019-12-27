package solvers

import (
	"fmt"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
)

func checkValid(board boards.Board, index int) bool {
	 //check horizontal
	 //check vertical
	 //check 3x3 squares
	return true
}

func checkHorizonal(board boards.Board, index int) bool {
	return true
}

func checkVertical(board boards.Board, index int) bool {
	return true
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
