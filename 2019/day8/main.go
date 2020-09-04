package main

import (
	"fmt"
	"github.com/mlhoyt/adventofcode.com-2019/day8/pkg/sifimage"
	"io/ioutil"
	"os"
)

var imageLayerWidth = 25
var imageLayerHeight = 6
var imageDataFile = "image-data.txt"

func main() {
	imageDataBytes, err := ioutil.ReadFile(imageDataFile)
	if err != nil {
		fmt.Printf("[ERROR] failed to read image data file %q: %v\n", imageDataFile, err)
		os.Exit(1)
	}

	imageData := []int{}
	for _, v := range string(imageDataBytes) {
		if v >= '0' && v <= '9' {
			imageData = append(imageData, int(v-'0'))
		}
	}

	image, err := sifimage.NewSIFImage(imageLayerWidth, imageLayerHeight, imageData)
	if err != nil {
		fmt.Printf("[ERROR] failed to decode image data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("image checkSum: %v\n", image.CheckSum())

	fmt.Printf("image rendered:\n%s\n", image.Render())
}
