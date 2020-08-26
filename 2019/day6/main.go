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

	// checkSum := uom.CheckSum()
	// fmt.Printf("checkSum %d\n", checkSum)

	youOrbits := uom.GetOrbiteeList("YOU")
	sanOrbits := uom.GetOrbiteeList("SAN")
	commonOrbitee := youOrbits.CalculateDivergencePoint(sanOrbits)

	youOrbitalDist, err := uom.GetOrbitalDistance("YOU", commonOrbitee)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		os.Exit(1)
	}
	sanOrbitalDist, err := uom.GetOrbitalDistance("SAN", commonOrbitee)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		os.Exit(1)
	}
	orbitalTransfers := youOrbitalDist + sanOrbitalDist
	fmt.Printf("orbital-transfers %d\n", orbitalTransfers)
}
