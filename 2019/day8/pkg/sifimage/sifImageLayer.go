package sifimage

import (
	"bytes"
	"fmt"
	"strconv"
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

func (u *sifImageLayer) merge(v *sifImageLayer) *sifImageLayer {
	merged := &sifImageLayer{
		data: [][]int{},
	}

	for rowIndex := range u.data {
		mergedRow := []int{}
		for columnIndex := range u.data[rowIndex] {
			mergedRow = append(mergedRow, mergeValues(u.data[rowIndex][columnIndex], v.data[rowIndex][columnIndex]))
		}
		merged.data = append(merged.data, mergedRow)
	}

	return merged
}

func mergeValues(v1 int, v2 int) int {
	if v1 == 2 {
		return v2
	}

	return v1
}

func (u *sifImageLayer) String() string {
	buf := &bytes.Buffer{}

	for _, row := range u.data {
		for _, rcv := range row {
			buf.WriteString(strconv.Itoa(rcv))
		}
		buf.WriteString("\n")
	}

	return buf.String()
}
