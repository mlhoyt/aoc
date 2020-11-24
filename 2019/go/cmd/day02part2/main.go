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

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			testCode := make([]int, len(code))
			copy(testCode, code)

			// Restore to before "1202 program alarm"
			testCode[1] = noun
			testCode[2] = verb

			computer := intcode.NewIntCode(testCode)

			if err := computer.Run(); err != nil {
				continue
			}

			if computer.MemAt(0) == 19690720 {
				fmt.Printf("%d\n", 100*noun+verb)
				return
			}
		}
	}

	fmt.Println("no noun/verb match found")
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
