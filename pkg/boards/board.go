package boards

import (
	"errors"
	"sudokusolver/pkg/cells"
)

const CELL_COUNT = 81
const BOARD_SIDE_LENGTH = 9

type boardStruct struct {
	complete     bool
	cellsFactory cells.Factory
	cells        []cells.IndividualCell
}
type Board interface {
	GetCell(index int) cells.IndividualCell
	SetCellValue(index int, value int)
	GetCells() *[]cells.IndividualCell
	SetComplete()
	CheckComplete() bool
}

func NewBoard(cellFactory cells.Factory, initialValues []int) (Board, error) {

	if len(initialValues) != CELL_COUNT {
		return nil, errors.New("Invalid number of sudoku values provided")
	}

	board := boardStruct{false, cellFactory, []cells.IndividualCell{}}
	for _, initialValue := range initialValues {
		newCell, err := board.cellsFactory.NewCell(initialValue)
		if(err != nil) {
			return nil, err
		}
		board.cells = append(board.cells, newCell)
	}
	return &board, nil
}

func (b boardStruct) GetCell(index int) cells.IndividualCell {
	if index < 0 || index > CELL_COUNT {
		panic("Invalid index used in GetCell()")
	}
	return b.cells[index]
}

func (b *boardStruct) SetCellValue(index int, newValue int) {
	cells.ValidateCellValue(newValue)
	if b.GetCell(index).GetCellType() == cells.SETTABLE_CELL_TYPE {
		b.cells[index].SetCellValue(newValue)
	} else {
		panic("Attempting to set a value to a cell that already has a preset value")
	}
}

func (b boardStruct) GetCells() *[]cells.IndividualCell {
	return &b.cells
}

func (b *boardStruct) SetComplete() {
	b.complete = true
}

func (b boardStruct) CheckComplete() bool {
	return b.complete
}
