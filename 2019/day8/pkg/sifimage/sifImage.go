package sifimage

import (
	"fmt"
)

var errInvalidImageDataLength = fmt.Errorf("invalid image data length; must be an even multiple of width times height")

type sifImage struct {
	layers []*sifImageLayer
}

func NewSIFImage(width int, height int, data []int) (*sifImage, error) {
	layerSize := width * height

	if len(data)%layerSize != 0 {
		return nil, errInvalidImageDataLength
	}

	layerCount := len(data) / layerSize

	image := &sifImage{
		layers: make([]*sifImageLayer, layerCount),
	}

	for layerNr := 0; layerNr < layerCount; layerNr++ {
		layerStart := layerNr * layerSize
		layerEnd := layerStart + layerSize

		layer, err := newSIFImageLayer(width, height, data[layerStart:layerEnd])
		if err != nil {
			return nil, err
		}

		image.layers[layerNr] = layer
	}

	return image, nil
}

func (u *sifImage) CheckSum() int {
	checkSumLayerNr := 0
	checkSumLayerZeroCount := u.layers[0].size() + 1
	for layerNr, layer := range u.layers {
		layerZeroCount := layer.valueCount(0)

		if layerZeroCount < checkSumLayerZeroCount {
			checkSumLayerNr = layerNr
			checkSumLayerZeroCount = layerZeroCount
		}
	}

	return u.layers[checkSumLayerNr].valueCount(1) * u.layers[checkSumLayerNr].valueCount(2)
}
