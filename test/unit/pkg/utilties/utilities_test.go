package utilties

import (
	"github.com/stretchr/testify/assert"
	utilties "sudokusolver/pkg/utilites"
	"testing"
)

func TestConvertStringSliceToIntSlice(t *testing.T) {
	original := []string{"1", "2", "3"}
	expected := []int{1, 2, 3}
	result := utilties.ConvertStringSliceToIntSlice(original)
	assert.Equal(t, expected, result)

	original = []string{"-1", "-2", "-3"}
	expected = []int{-1, -2, -3}
	result = utilties.ConvertStringSliceToIntSlice(original)
	assert.Equal(t, expected, result)

	assert.Panics(t, func() {
		original = []string{"a", "b", "c"}
		expected = []int{}
		result = utilties.ConvertStringSliceToIntSlice(original)
	})

	assert.Panics(t, func() {
		original = []string{"arg1", "€", "驰馳"}
		expected = []int{}
		result = utilties.ConvertStringSliceToIntSlice(original)
	})

	assert.Panics(t, func() {
		original = []string{""}
		expected = []int{}
		result = utilties.ConvertStringSliceToIntSlice(original)
	})
}

func TestStringInSlice(t *testing.T) {

	searchItems := []string{"arg1", "€", "驰馳", ""}

	assert.True(t, utilties.StringIntSlice("arg1", searchItems))
	assert.True(t, utilties.StringIntSlice("€", searchItems))
	assert.True(t, utilties.StringIntSlice("驰馳", searchItems))
	assert.True(t, utilties.StringIntSlice("", searchItems))
	assert.False(t, utilties.StringIntSlice("0", searchItems))
	assert.False(t, utilties.StringIntSlice("NotHere", searchItems))
}

func TestIntInSlice(t *testing.T) {

	searchItems := []int{0, 1, 1313242523524346, -123123}

	assert.True(t, utilties.IntInSlice(0, searchItems))
	assert.True(t, utilties.IntInSlice(1, searchItems))
	assert.True(t, utilties.IntInSlice(1313242523524346, searchItems))
	assert.True(t, utilties.IntInSlice(-123123, searchItems))
	assert.False(t, utilties.IntInSlice(98, searchItems))
	assert.False(t, utilties.IntInSlice(89453843573458, searchItems))
	assert.False(t, utilties.IntInSlice(-89453843573458, searchItems))
}

func TestAreSudokuValuesUnique(t *testing.T) {
	uniqueItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.True(t, utilties.AreSudokuValuesUnique(uniqueItems))

	uniqueItems2 := []int{1}
	assert.True(t, utilties.AreSudokuValuesUnique(uniqueItems2))

	uniqueItems3 := []int{0}
	assert.True(t, utilties.AreSudokuValuesUnique(uniqueItems3))

	nonUniqueItems := []int{2, 2, 2, 4, 5, 6, 7, 8, 9}
	assert.False(t, utilties.AreSudokuValuesUnique(nonUniqueItems))

	nonUniqueItems2 := []int{9, 7, 6, 6, 5, 4, 3, 3, 2, 2}
	assert.False(t, utilties.AreSudokuValuesUnique(nonUniqueItems2))

	//We disregard zeros as we're not comparing them, so multiple zeros
	// Will not be evaluated as being non-unique
	nonUniqueItems3 := []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	assert.True(t, utilties.AreSudokuValuesUnique(nonUniqueItems3))
}
