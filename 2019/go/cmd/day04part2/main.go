package main

import (
	"fmt"
	"github.com/mlhoyt/advent-of-code/2019/go/pkg/utils"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.LoadInputFile("day04.txt")
	if err != nil {
		panic(err)
	}

	min, max, err := newBounds(input)
	if err != nil {
		panic(err)
	}

	nrValidPasswords := calculateValidPasswordsInRange(min, max)

	fmt.Printf("%d\n", nrValidPasswords)
}

func newBounds(data []string) (int, int, error) {
	bounds := strings.Split(data[0], "-")

	min, err := strconv.ParseInt(bounds[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}

	max, err := strconv.ParseInt(bounds[1], 10, 32)
	if err != nil {
		return 0, 0, err
	}

	return int(min), int(max), nil
}

func calculateValidPasswordsInRange(min int, max int) int {
	nrValidPasswords := 0
	for _, password := range generatePasswords(min, max) {
		if isValidPassword(password, hasAdjacentMatch, hasAdjacentMatchLength(2), isMonotonicallyIncreasing) {
			nrValidPasswords++
		}
	}

	return nrValidPasswords
}

func generatePasswords(min int, max int) []string {
	passwords := []string{}
	for n := min; n <= max; n++ {
		nStr := strconv.FormatInt(int64(n), 10)
		passwords = append(passwords, nStr)
	}

	return passwords
}

func isValidPassword(v string, ps ...func(string) bool) bool {
	for _, p := range ps {
		if !p(v) {
			return false
		}
	}

	return true
}

func hasAdjacentMatch(v string) bool {
	for i := 1; i < len(v); i++ {
		if v[i] == v[i-1] {
			return true
		}
	}

	return false
}

func hasAdjacentMatchLength(n int) func(string) bool {
	p := func(v string) bool {
		matchLength := 1
		for i := 1; i < len(v); i++ {
			if v[i] == v[i-1] {
				matchLength++
			} else {
				if matchLength == n {
					return true
				}
				matchLength = 1
			}
		}

		return matchLength == n
	}

	return p
}

func isMonotonicallyIncreasing(v string) bool {
	for i := 1; i < len(v); i++ {
		if v[i] < v[i-1] {
			return false
		}
	}

	return true
}
