package main

import (
	"fmt"
	"math"
)

var modules = []int{
	51585,
	137484,
	73634,
	71535,
	87274,
	74243,
	127025,
	66829,
	138729,
	145459,
	118813,
	82326,
	82518,
	145032,
	148699,
	105958,
	103969,
	72689,
	145061,
	70385,
	53104,
	107851,
	103392,
	107051,
	123475,
	123918,
	56709,
	89284,
	86208,
	71943,
	109257,
	108272,
	124811,
	142709,
	115650,
	53607,
	142891,
	144135,
	114277,
	138671,
	111998,
	70838,
	69802,
	107210,
	103319,
	60377,
	58639,
	131863,
	100807,
	118360,
	52573,
	108207,
	128009,
	96180,
	148492,
	112914,
	72867,
	140991,
	131267,
	125123,
	58393,
	129615,
	87239,
	63085,
	59231,
	95007,
	147712,
	109838,
	89829,
	55634,
	96163,
	52323,
	106701,
	141511,
	125349,
	137267,
	50694,
	53692,
	57466,
	117769,
	63535,
	101708,
	113593,
	79163,
	112327,
	91994,
	129674,
	58076,
	145062,
	122730,
	102481,
	109994,
	136271,
	111178,
	117920,
	107933,
	104305,
	99613,
	68482,
	126543,
}

// var modules = []int{
// 	12,
// 	14,
// 	1969,
// 	100756,
// }

// Part 1
// func massToFuel(mass int) int {
// 	return int(math.Floor(float64(mass)/3.0)) - 2
// }

// Part 2
func massToFuel(mass int) int {
	if mass <= 0 {
		return 0
	}

	fuel := int(math.Floor(float64(mass)/3.0)) - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + massToFuel(fuel)
}

func main() {
	totalFuel := 0
	for _, moduleMass := range modules {
		// moduleFuel := massToFuel(moduleMass)
		// fmt.Printf("moduleMass=%d moduleFuel=%d\n", moduleMass, moduleFuel)

		// totalFuel = totalFuel + moduleFuel
		totalFuel = totalFuel + massToFuel(moduleMass)
	}

	fmt.Printf("totalFuel=%d\n", totalFuel)
}
