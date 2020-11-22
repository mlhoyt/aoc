package main

import (
	"bufio"
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/wireanalyzer"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	input, err := load_file("input/day03.txt")
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
