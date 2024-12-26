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
	orbitalDistance := youOrbits.CalculateOrbitalDistance(sanOrbits)
	fmt.Printf("orbitalDistance: %d\n", orbitalDistance)
}
