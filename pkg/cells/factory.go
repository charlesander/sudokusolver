package cells

import utilties "sudokusolver/pkg/utilites"

type fact struct {
}

//NewService Creates a new output service
func NewFactory() Factory {
	return &fact{}
}

//Service Define the output service interface
type Factory interface {
	NewCell(CellType string, CellValue int) IndividualCell
}

func (f fact) NewCell(CellType string, CellValue int) IndividualCell {
	validCellTypes := []string{"Preset", "Settable"}
	if !utilties.StringInSlice(CellType, validCellTypes) {
		panic("Invalid CellType provided")
	}
	if CellValue < 0 || CellValue > 9 {
		panic("Invalid CellValue provided")
	}
	return &individual{CellType, CellValue}
}