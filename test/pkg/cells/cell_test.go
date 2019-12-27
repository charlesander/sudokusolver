package cells

import (
	"github.com/stretchr/testify/assert"
	"sudokusolver/pkg/cells"
	"testing"
)

func TestSum(t *testing.T) {
	cell := cells.NewPreset("Preset", 1);
	assert.Equal(t, cell.GetCellType(), "Preset")
	assert.Equal(t, cell.GetCellValue(), 1)

	cell2 := cells.NewPreset("Settable", 9);
	assert.Equal(t, cell2.GetCellType(), "Settable")
	assert.Equal(t, cell2.GetCellValue(), 9)

	assert.Panics(t, func() {
		cells.NewPreset("Invalid", 1)
	})

	assert.Panics(t, func() {
		cells.NewPreset("Invalid", -1)
	})

	assert.Panics(t, func() {
		cells.NewPreset("Invalid", 10)
	})

	assert.Panics(t, func() {
		cells.NewPreset("Preset", -500)
	})

	assert.Panics(t, func() {
		cells.NewPreset("Preset", 123)
	})

	assert.Panics(t, func() {
		cells.NewPreset("Settable", -1)
	})

	assert.Panics(t, func() {
		cells.NewPreset("Settable", 10)
	})
}