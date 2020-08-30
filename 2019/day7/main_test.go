package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulateAmpChain1(t *testing.T) {
	code := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	setting := []int{4, 3, 2, 1, 0}

	output, err := simulateAmpChain(code, setting)
	assert.Nil(t, err)
	assert.Equal(t, 43210, output)
}
