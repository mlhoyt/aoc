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

func TestSIFImageLayerMergeValues(t *testing.T) {
	tests := []struct {
		name     string
		v1       int
		v2       int
		expected int
	}{
		{
			name:     "0, 0 = 0",
			v1:       0,
			v2:       0,
			expected: 0,
		},
		{
			name:     "0, 1 = 0",
			v1:       0,
			v2:       1,
			expected: 0,
		},
		{
			name:     "0, 2 = 0",
			v1:       0,
			v2:       2,
			expected: 0,
		},
		{
			name:     "1, 0 = 1",
			v1:       1,
			v2:       0,
			expected: 1,
		},
		{
			name:     "1, 1 = 1",
			v1:       1,
			v2:       1,
			expected: 1,
		},
		{
			name:     "1, 2 = 1",
			v1:       1,
			v2:       2,
			expected: 1,
		},
		{
			name:     "2, 0 = 0",
			v1:       2,
			v2:       0,
			expected: 0,
		},
		{
			name:     "2, 1 = 1",
			v1:       2,
			v2:       1,
			expected: 1,
		},
		{
			name:     "2, 2 = 2",
			v1:       2,
			v2:       2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mergeValues(tt.v1, tt.v2)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewSIFImageLayerString(t *testing.T) {
	layer, err := newSIFImageLayer(2, 2, []int{0, 2, 1, 3})
	expected := "02\n13\n"

	assert.Nil(t, err)
	assert.Equal(t, expected, layer.String())
}

func TestNewSIFImageLayerMerge(t *testing.T) {
	layer1, err1 := newSIFImageLayer(2, 2, []int{0, 1, 2, 0})
	layer2, err2 := newSIFImageLayer(2, 2, []int{1, 2, 0, 1})
	expected := "01\n00\n"

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, expected, layer1.merge(layer2).String())
}
