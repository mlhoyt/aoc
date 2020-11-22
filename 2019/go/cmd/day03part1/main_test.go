package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNearestWireIntersection1(t *testing.T) {
	w1 := []string{"R8", "U5", "L5", "D3"}
	w2 := []string{"U7", "R6", "D4", "L4"}
	expected := 6

	actual := nearestWireIntersection(w1, w2)
	assert.Equal(t, expected, actual)
}

func TestNearestWireIntersection2(t *testing.T) {
	w1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	w2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	expected := 159

	actual := nearestWireIntersection(w1, w2)
	assert.Equal(t, expected, actual)
}

func TestNearestWireIntersection3(t *testing.T) {
	w1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	w2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	expected := 135

	actual := nearestWireIntersection(w1, w2)
	assert.Equal(t, expected, actual)
}

// func TestFewestStepsWireIntersection1(t *testing.T) {
// 	w1 := []string{"R8", "U5", "L5", "D3"}
// 	w2 := []string{"U7", "R6", "D4", "L4"}
// 	expected := 30

// 	actual := FewestStepsWireIntersection(w1, w2)
// 	assert.Equal(t, expected, actual)
// }

// func TestFewestStepsWireIntersection2(t *testing.T) {
// 	w1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
// 	w2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
// 	expected := 610

// 	actual := FewestStepsWireIntersection(w1, w2)
// 	assert.Equal(t, expected, actual)
// }

// func TestFewestStepsWireIntersection3(t *testing.T) {
// 	w1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
// 	w2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
// 	expected := 410

// 	actual := FewestStepsWireIntersection(w1, w2)
// 	assert.Equal(t, expected, actual)
// }
