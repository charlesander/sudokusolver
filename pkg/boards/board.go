package board

import (
	"sudokusolver/pkg/cells"
)

type boardStruct struct {
	cellsFactory cells.Factory
	cells        []cells.IndividualCell
}
type Board interface {
	GetCell(index int) cells.IndividualCell
	SetCellValue(index int, value int)
	GetCells() *[]cells.IndividualCell
	CheckComplete() bool
}

func NewBoard(cellFactory cells.Factory, initialValues []int) Board {

	if len(initialValues) != 81 {
		panic("Invalid number of sudoku values provided")
	}

	board := boardStruct{cellFactory, []cells.IndividualCell{}}
	for _, initialValue := range initialValues {
		board.cells = append(board.cells, board.cellsFactory.NewCell(initialValue))
	}
	return board
}

func (b boardStruct) GetCell(index int) cells.IndividualCell {
	return b.cells[index]
}

func (b boardStruct) SetCellValue(index int, newValue int) {
	cells.ValidateCellValue(newValue)
	if(b.GetCell(index).GetCellType() == cells.SETTABLE_CELL_TYPE) {
		b.cells[index].SetCellValue(newValue)
	} else {
		panic("Attempting to set a value to a cell that already has a preset value")
	}
}

func (b boardStruct) GetCells() *[]cells.IndividualCell {
	return &b.cells
}

func (b boardStruct) CheckComplete() bool {
	return true
}