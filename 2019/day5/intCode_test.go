package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunOpAddIndirect(t *testing.T) {
	code := []int{1, 2, 3, 5, 99, 0}
	expected := []int{1, 2, 3, 5, 99, 8}

	intcode := NewIntCode(code, os.Stdin, os.Stdout)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpMultIndirect(t *testing.T) {
	code := []int{2, 2, 3, 5, 99, 0}
	expected := []int{2, 2, 3, 5, 99, 15}

	intcode := NewIntCode(code, os.Stdin, os.Stdout)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpAddIndirectMultIndirect(t *testing.T) {
	code := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	expected := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}

	intcode := NewIntCode(code, os.Stdin, os.Stdout)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpInput(t *testing.T) {
	code := []int{3, 1, 99}
	expected := []int{3, 13, 99}

	var stdin bytes.Buffer
	stdin.Write([]byte("13\n"))

	intcode := NewIntCode(code, &stdin, os.Stdout)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpOutput(t *testing.T) {
	code := []int{4, 3, 99, 13}
	expected := []int{4, 3, 99, 13}

	var stdout bytes.Buffer

	intcode := NewIntCode(code, os.Stdin, &stdout)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
	assert.Equal(t, []byte("13\n"), stdout.Bytes())
}
