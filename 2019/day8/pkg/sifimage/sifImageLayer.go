package sifimage

import (
	"fmt"
)

var errInvalidImageLayerDataLength = fmt.Errorf("invalid image layer data length; must be an even multiple of width times height")

type sifImageLayer struct {
	data [][]int
}

func newSIFImageLayer(width int, height int, data []int) (*sifImageLayer, error) {
	layerSize := width * height

	if len(data)%layerSize != 0 {
		return nil, errInvalidImageLayerDataLength
	}

	rowCount := len(data) / width

	layer := &sifImageLayer{
		data: make([][]int, rowCount),
	}

	for rowNr := 0; rowNr < rowCount; rowNr++ {
		rowStart := rowNr * width
		rowEnd := rowStart + width

		row := make([]int, width)
		copy(row, data[rowStart:rowEnd])

		layer.data[rowNr] = row
	}

	return layer, nil
}

func (u *sifImageLayer) size() int {
	return len(u.data) * len(u.data[0])
}

func (u *sifImageLayer) valueCount(v int) int {
	count := 0
	for _, row := range u.data {
		for _, rcv := range row {
			if rcv == v {
				count++
			}
		}
	}

	return count
}
