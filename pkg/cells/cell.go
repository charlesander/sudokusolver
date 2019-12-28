package cells

const SETTABLE_CELL_TYPE = "Settable"
const PRESET_CELL_TYPE = "Preset"
const MIN_CELL_VALUE = 0
const MAX_CELL_VALUE = 9

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

func (p *individual) SetCellValue(newValue int) {
	p.cellValue = newValue
}

func ValidateCellValue(CellValue int) {
	if CellValue < MIN_CELL_VALUE || CellValue > MAX_CELL_VALUE {
		panic("Invalid CellValue provided")
	}
}