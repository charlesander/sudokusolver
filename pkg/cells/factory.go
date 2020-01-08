package cells

type fact struct {
}

//NewService Creates a new output service
func NewFactory() Factory {
	return &fact{}
}

//Service Define the output service interface
type Factory interface {
	NewCell(CellValue int) (IndividualCell, error)
}

func (f fact) NewCell(CellValue int) (IndividualCell, error) {
	err := ValidateCellValue(CellValue)
	if(err != nil) {
		return nil, err
	}
	var CellType string
	if CellValue == 0 {
		CellType = SETTABLE_CELL_TYPE
	} else {
		CellType = PRESET_CELL_TYPE
	}
	return &individual{CellType, CellValue}, nil
}