package main

import (
	"fmt"
	"os"
)

func main() {
	uom, err := NewUniversalOrbitMapFromFile("input.txt")
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		os.Exit(1)
	}

	youOrbits := uom.GetOrbiteeList("YOU")
	sanOrbits := uom.GetOrbiteeList("SAN")
	commonOrbiteeIndex := youOrbits.CalculateDivergencePoint(sanOrbits)

	youOrbitalDistance := youOrbits.CalculateOrbitalDistance(commonOrbiteeIndex)
	sanOrbitalDistance := sanOrbits.CalculateOrbitalDistance(commonOrbiteeIndex)

	orbitalTransfers := youOrbitalDistance + sanOrbitalDistance
	fmt.Printf("orbital-transfers %d\n", orbitalTransfers)
}
