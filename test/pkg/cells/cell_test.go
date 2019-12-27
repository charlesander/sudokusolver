package cells

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/cells"
	"testing"
)

func TestCellInitiation(t *testing.T) {

	cell := cells.NewCell("Preset", 1);
	assert.Equal(t, cell.GetCellType(), "Preset")
	assert.Equal(t, cell.GetCellValue(), 1)

	cell2 := cells.NewCell("Settable", 9);
	assert.Equal(t, cell2.GetCellType(), "Settable")
	assert.Equal(t, cell2.GetCellValue(), 9)
}

func TestCellInitiationWithInvalidData(t *testing.T) {
	assert.Panics(t, func() {
		cells.NewCell("Invalid", 1)
	})

	assert.Panics(t, func() {
		cells.NewCell("Invalid", -1)
	})

	assert.Panics(t, func() {
		cells.NewCell("Invalid", 10)
	})

	assert.Panics(t, func() {
		cells.NewCell("Preset", -500)
	})

	assert.Panics(t, func() {
		cells.NewCell("Preset", 123)
	})

	assert.Panics(t, func() {
		cells.NewCell("Settable", -1)
	})

	assert.Panics(t, func() {
		cells.NewCell("Settable", 10)
	})
}