package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunOpAddIndirect(t *testing.T) {
	code := []int{1, 2, 3, 5, 99, 0}
	expected := []int{1, 2, 3, 5, 99, 8}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpMultIndirect(t *testing.T) {
	code := []int{2, 2, 3, 5, 99, 0}
	expected := []int{2, 2, 3, 5, 99, 15}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpAddIndirectMultIndirect(t *testing.T) {
	code := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	expected := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpInput(t *testing.T) {
	t.Skip() // FIXME: need to control stdin and pass 13
	code := []int{3, 1, 99}
	expected := []int{3, 13, 99}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpOutput(t *testing.T) {
	t.Skip() // FIXME: need to control stdout and read 13
	code := []int{3, 3, 99, 13}
	expected := []int{3, 3, 99, 13}

	intcode := NewIntCode(code)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}
