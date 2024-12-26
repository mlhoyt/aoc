package sifimage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSIFImageInvalidDataLength(t *testing.T) {
	image, err := NewSIFImage(2, 2, []int{0, 1, 2, 3, 4, 5})

	assert.Equal(t, errInvalidImageDataLength, err)
	assert.Nil(t, image)
}

func TestNewSIFImage(t *testing.T) {
	image, err := NewSIFImage(2, 2, []int{0, 2, 1, 3, 10, 12, 11, 13})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(image.layers))
	assert.Equal(t, [][]int{[]int{0, 2}, []int{1, 3}}, image.layers[0].data)
	assert.Equal(t, [][]int{[]int{10, 12}, []int{11, 13}}, image.layers[1].data)
}

func TestSIFImageCheckSum(t *testing.T) {
	image, err := NewSIFImage(2, 2, []int{0, 2, 1, 3, 1, 1, 2, 2})

	assert.Nil(t, err)
	assert.Equal(t, 4, image.CheckSum())
}
