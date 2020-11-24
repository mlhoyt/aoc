package main

import (
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/utils"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/wireanalyzer"
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

	steps := fewestStepsWireIntersection(wires[0], wires[1])
	fmt.Printf("%d\n", steps)
}

func newWires(data []string) ([]wireanalyzer.Wire, error) {
	newWires := []wireanalyzer.Wire{}
	for _, v := range data {
		newWires = append(newWires, wireanalyzer.Wire(strings.Split(v, ",")))
	}

	return newWires, nil
}

func fewestStepsWireIntersection(w1 wireanalyzer.Wire, w2 wireanalyzer.Wire) int {
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

	minSteps := 0
	for _, intersection := range intersections {
		p1Steps := p1.StepsToIntersection(intersection)
		p2Steps := p2.StepsToIntersection(intersection)
		steps := p1Steps + p2Steps
		if minSteps == 0 || steps < minSteps {
			minSteps = steps
		}
	}

	return minSteps
}
