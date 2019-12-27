package cells

type individual struct {
	cellType  string
	cellValue int
}

type IndividualCell interface {
	GetCellType() string
	GetCellValue() int
	SetCellValue(newValue int)
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
