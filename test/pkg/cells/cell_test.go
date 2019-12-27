package cells

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/cells"
	"testing"
)

func TestCellInitiation(t *testing.T) {
	cellFactory := cells.NewFactory()
	cell := cellFactory.NewCell("Preset", 1);
	assert.Equal(t, cell.GetCellType(), "Preset")
	assert.Equal(t, cell.GetCellValue(), 1)

	cell2 := cellFactory.NewCell("Settable", 9);
	assert.Equal(t, cell2.GetCellType(), "Settable")
	assert.Equal(t, cell2.GetCellValue(), 9)
}

func TestCellInitiationWithInvalidData(t *testing.T) {
	cellFactory := cells.NewFactory()
	assert.Panics(t, func() {
		cellFactory.NewCell("Invalid", 1)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell("Invalid", -1)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell("Invalid", 10)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell("Preset", -500)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell("Preset", 123)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell("Settable", -1)
	})

	assert.Panics(t, func() {
		cellFactory.NewCell("Settable", 10)
	})
}