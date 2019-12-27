package utilties

import (
	"github.com/stretchr/testify/assert"
	utilties "sudokusolver/pkg/utilites"
	"testing"
)

func TestSum(t *testing.T) {

	searchItems := []string{"arg1", "€", "驰馳", ""}

	assert.True(t, utilties.StringInSlice("arg1", searchItems))
	assert.True(t, utilties.StringInSlice("€", searchItems))
	assert.True(t, utilties.StringInSlice("驰馳", searchItems))
	assert.True(t, utilties.StringInSlice("", searchItems))
	assert.False(t, utilties.StringInSlice("0", searchItems))
	assert.False(t, utilties.StringInSlice("NotHere", searchItems))
}