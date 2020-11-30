package main

import (
	"bytes"
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/intcode"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/utils"
	"os"
	"strconv"
)

var stdinData = []byte("1\n")

// part #2
// var stdinData = []byte("5\n")

func main() {
	input, err := utils.LoadInputFile("day05.txt")
	if err != nil {
		panic(err)
	}

	code, err := newCode(input)
	if err != nil {
		panic(err)
	}

	stdin := bytes.NewBuffer(stdinData)
	stdout := bytes.Buffer{}

	intcode := intcode.NewIntCode(code, stdin, &stdout)

	if err := intcode.Run(); err != nil {
		fmt.Printf("[ERROR] runtime error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(stdout.String())
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
