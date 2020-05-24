package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunOpCode1(t *testing.T) {
	code := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpCode2(t *testing.T) {
	code := []int{2, 3, 0, 3, 99}
	expected := []int{2, 3, 0, 6, 99}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpCode1Plus(t *testing.T) {
	code := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpCode2Plus(t *testing.T) {
	code := []int{2, 4, 4, 5, 99, 0}
	expected := []int{2, 4, 4, 5, 99, 9801}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpCode12(t *testing.T) {
	code := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	expected := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}
