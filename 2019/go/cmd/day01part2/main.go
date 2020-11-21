package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	input, err := load_file("input/day01.txt")
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

func load_file(name string) ([]string, error) {
	absName, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}

	ifh, err := os.Open(absName)
	if err != nil {
		return nil, err
	}
	defer ifh.Close()

	lines := []string{}
	scanner := bufio.NewScanner(ifh)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
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
		return 0
	}

	return fuel + massToFuel(fuel)
}
