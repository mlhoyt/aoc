package main

import (
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/intcode"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/utils"
	"strconv"
)

func main() {
	input, err := utils.LoadInputFile("day02.txt")
	if err != nil {
		panic(err)
	}

	code, err := newCode(input)
	if err != nil {
		panic(err)
	}

	// Restore to before "1202 program alarm"
	code[1] = 12
	code[2] = 2

	computer := intcode.NewIntCode(code, nil, nil)

	if err := computer.Run(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", computer.MemAt(0))
}

func newCode(data []string) ([]int, error) {
	newCode := make([]int, len(data))
	for i, v := range data {
		vInt64, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return newCode, err
		}

		newCode[i] = int(vInt64)
	}

	return newCode, nil
}
