package cells

type fact struct {
}

//NewService Creates a new output service
func NewFactory() Factory {
	return &fact{}
}

//Service Define the output service interface
type Factory interface {
	NewCell(CellValue int) IndividualCell
}

func (f fact) NewCell(CellValue int) IndividualCell {
	if CellValue < 0 || CellValue > 9 {
		panic("Invalid CellValue provided")
	}
	var CellType string
	if CellValue == 0 {
		CellType = "Settable"
	} else {
		CellType = "Preset"
	}
	return &individual{CellType, CellValue}
}