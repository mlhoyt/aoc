package main

import (
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/utils"
	"math"
	"strconv"
)

func main() {
	input, err := utils.LoadInputFile("day01.txt")
	if err != nil {
		panic(err)
	}

	modules, err := newModules(input)
	if err != nil {
		panic(err)
	}

	fuel := modules.calculateFuel(massToFuel)

	fmt.Printf("%d\n", fuel)
}

type modules []int

func newModules(data []string) (modules, error) {
	newModules := make(modules, len(data))
	for i, v := range data {
		vInt64, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return newModules, err
		}

		newModules[i] = int(vInt64)
	}

	return newModules, nil
}

func (u modules) calculateFuel(f func(int) int) int {
	fuel := 0
	for _, module := range u {
		fuel += f(module)
	}

	return fuel
}

func massToFuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3.0)) - 2
	if fuel < 0 {
		fuel = 0
	}

	return fuel
}
