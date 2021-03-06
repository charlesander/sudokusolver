package cells

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/cells"
	"testing"
)

func TestCellInitiation(t *testing.T) {
	cellFactory := cells.NewFactory()
	cell := cellFactory.NewCell( 1);
	assert.Equal(t, cell.GetCellType(), "Preset")
	assert.Equal(t, cell.GetCellValue(), 1)

	cell2 := cellFactory.NewCell(9);
	assert.Equal(t, cell2.GetCellType(), "Preset")
	assert.Equal(t, cell2.GetCellValue(), 9)

	cell3 := cellFactory.NewCell(0);
	assert.Equal(t, cell3.GetCellType(), "Settable")
	assert.Equal(t, cell3.GetCellValue(), 0)
}

func TestSetCellValue(t *testing.T) {
	cellFactory := cells.NewFactory()
	cell := cellFactory.NewCell( 1);
	cell.SetCellValue(2)
	assert.Equal(t, cell.GetCellValue(), 2)
}

func TestCellInitiationWithInvalidData(t *testing.T) {
	cellFactory := cells.NewFactory()

	assert.Panics(t, func() {
		cellFactory.NewCell( -1)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell(10)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell( -500)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell(123)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell( -1)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell(10)
	})
}

func TestValidateCellValue(t *testing.T) {
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(1)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(2)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(3)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(4)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(5)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(6)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(7)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(8)
	})
	assert.NotPanics(t, func() {
		cells.ValidateCellValue(9)
	})

	assert.Panics(t, func() {
		cells.ValidateCellValue(-1)
	})
	assert.Panics(t, func() {
		cells.ValidateCellValue(-10)
	})
	assert.Panics(t, func() {
		cells.ValidateCellValue(-123423434234)
	})
	assert.Panics(t, func() {
		cells.ValidateCellValue(10)
	})
	assert.Panics(t, func() {
		cells.ValidateCellValue(132423430)
	})

}