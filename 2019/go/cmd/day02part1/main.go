package main

import (
	"bufio"
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/intcode"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	input, err := load_file("input/day02.txt")
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

	computer := intcode.NewIntCode(code)

	if err := computer.Run(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", computer.MemAt(0))
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
