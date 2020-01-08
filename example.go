package main

import (
	"fmt"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
)

func mainExample() {
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

	solvedEasyBoard, err := boards.NewBoard(cells.NewFactory(), easySudoku)
	if err != nil {
		//TODO
	}
	solvers.Solve(solvedEasyBoard)
	fmt.Println(solvedEasyBoard.CheckComplete())

	echoBoard(solvedEasyBoard)

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

	solvedHardBoard, err := boards.NewBoard(cells.NewFactory(), hardSudoku)
	if err != nil {
		//TODO
	}
	solvers.Solve(solvedHardBoard)
	fmt.Println(solvedHardBoard.CheckComplete())

	echoBoard(solvedHardBoard)
}

func echoBoard(solvedBoard boards.Board) error{
	for i := 0; i < boards.CELL_COUNT; i++ {
		cell, err := solvedBoard.GetCell(i)
		if(err != nil) {
			return err
		}
		fmt.Print(cell.GetCellValue(), ",")
		//fmt.Print( i , ",")
		if (i+1)%9 == 0 {
			fmt.Print("\n")
		}
	}
	return nil
}
