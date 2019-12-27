package cells

import utilties "sudokusolver/pkg/utilites"

type individual struct {
	cellType  string
	cellValue int
}

type individualCell interface {
	GetCellType() string
	GetCellValue() int
	SetCellValue(newValue int)
}

func NewCell(CellType string, CellValue int) individualCell {
	validCellTypes := []string{"Preset", "Settable"}
	if !utilties.StringInSlice(CellType, validCellTypes) {
		panic("Invalid CellType provided")
	}
	if CellValue < 0 || CellValue > 9 {
		panic("Invalid CellValue provided")
	}
	return &individual{CellType, CellValue}
}

func (p individual) GetCellType() string {
	return p.cellType
}

func (p individual) GetCellValue() int {
	return p.cellValue
}

func (p individual) SetCellValue(newValue int) {
	p.cellValue = newValue
}
