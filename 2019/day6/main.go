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

	checkSum := uom.CheckSum()
	fmt.Printf("checkSum %d\n", checkSum)
}
