package main

import (
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/utils"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/wireanalyzer"
	"math"
	"strings"
)

func main() {
	input, err := utils.LoadInputFile("day03.txt")
	if err != nil {
		panic(err)
	}

	wires, err := newWires(input)
	if err != nil {
		panic(err)
	}
	if len(wires) < 2 {
		panic(fmt.Errorf("too few wires in"))
	}

	dist := nearestWireIntersection(wires[0], wires[1])
	fmt.Printf("%d\n", dist)
}

func newWires(data []string) ([]wireanalyzer.Wire, error) {
	newWires := []wireanalyzer.Wire{}
	for _, v := range data {
		newWires = append(newWires, wireanalyzer.Wire(strings.Split(v, ",")))
	}

	return newWires, nil
}

func nearestWireIntersection(w1 wireanalyzer.Wire, w2 wireanalyzer.Wire) int {
	p1, err := wireanalyzer.NewPathFromWire(w1)
	if err != nil {
		panic(err)
	}

	p2, err := wireanalyzer.NewPathFromWire(w2)
	if err != nil {
		panic(err)
	}

	intersections, err := p1.GetIntersections(p2)
	if err != nil {
		panic(err)
	}

	minDist := 0
	for _, intersection := range intersections {
		dist := int(math.Abs(float64(intersection.X))) + int(math.Abs(float64(intersection.Y)))
		if minDist == 0 || dist < minDist {
			minDist = dist
		}
	}

	return minDist
}
