package main

import (
	"fmt"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
)

func main() {
	fmt.Println("Sudoku solvers")

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

	solvedBoard := solvers.Solve(board)
	fmt.Println(solvedBoard.CheckComplete())

	for i := 0; i < boards.CELL_COUNT; i++ {
		fmt.Print(solvedBoard.GetCell(i).GetCellValue(), ",")
		//fmt.Print( i , ",")
		if (i+1)%9 == 0 {
			fmt.Print("\n")
		}
	}
}
