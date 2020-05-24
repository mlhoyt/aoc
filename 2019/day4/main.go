package main

import "fmt"
import "strconv"

func main() {
	min := 359282
	max := 820401

	nrValidPasswords := 0
	for n := min; n <= max; n++ {
		v := strconv.FormatInt(int64(n), 10)

		if isValidPassword(v) {
			fmt.Printf("[DEBUG] valid password: %s\n", v)
			nrValidPasswords++
		}
	}

	fmt.Printf("nrValidPasswords: %d\n", nrValidPasswords)
}

func isValidPassword(v string) bool {
	hasAdjacentMatch := false
	adjacentMatchLength := 1
	for i := 1; i < len(v); i++ {
		if v[i] == v[i-1] {
			adjacentMatchLength++
		} else {
			if adjacentMatchLength == 2 {
				hasAdjacentMatch = true
			}
			adjacentMatchLength = 1
		}

		if v[i] < v[i-1] {
			return false
		}
	}
	if adjacentMatchLength == 2 {
		hasAdjacentMatch = true
	}

	if !hasAdjacentMatch {
		return false
	}

	return true
}
