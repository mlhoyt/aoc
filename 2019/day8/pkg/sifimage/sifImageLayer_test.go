package sifimage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSIFImageLayerInvalidDataLength(t *testing.T) {
	layer, err := newSIFImageLayer(2, 2, []int{0, 1, 2})

	assert.Nil(t, layer)
	assert.Equal(t, errInvalidImageLayerDataLength, err)
}

func TestNewSIFImageLayer(t *testing.T) {
	layer, err := newSIFImageLayer(2, 2, []int{0, 2, 1, 3})
	expected := [][]int{
		[]int{0, 2},
		[]int{1, 3},
	}

	assert.Equal(t, expected, layer.data)
	assert.Nil(t, err)
}

func TestSIFImageLayerSize(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		data     []int
		expected int
	}{
		{
			name:     "2 x 2",
			width:    2,
			height:   2,
			data:     []int{0, 2, 1, 3},
			expected: 4,
		},
		{
			name:     "3 x 3",
			width:    3,
			height:   3,
			data:     []int{1, 3, 5, 2, 4, 6, 3, 5, 7},
			expected: 9,
		},
		{
			name:     "1 x 2",
			width:    1,
			height:   2,
			data:     []int{1, 3},
			expected: 2,
		},
		{
			name:     "2 x 1",
			width:    2,
			height:   1,
			data:     []int{1, 3},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layer, err := newSIFImageLayer(tt.width, tt.height, tt.data)

			assert.Nil(t, err)
			assert.Equal(t, tt.expected, layer.size())
		})
	}
}

func TestSIFImageLayerValueCount(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		data     []int
		value    int
		expected int
	}{
		{
			name:     "1 zero",
			width:    2,
			height:   2,
			data:     []int{0, 2, 1, 3},
			value:    0,
			expected: 1,
		},
		{
			name:     "2 threes",
			width:    3,
			height:   3,
			data:     []int{1, 3, 5, 2, 4, 6, 3, 5, 7},
			value:    3,
			expected: 2,
		},
		{
			name:     "0 twos",
			width:    1,
			height:   2,
			data:     []int{1, 3},
			value:    2,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layer, err := newSIFImageLayer(tt.width, tt.height, tt.data)

			assert.Nil(t, err)
			assert.Equal(t, tt.expected, layer.valueCount(tt.value))
		})
	}
}
