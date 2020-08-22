package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunOpAdd(t *testing.T) {
	tests := []struct {
		name     string
		code     []int
		expected []int
	}{
		{
			name:     "add indirect indirect",
			code:     []int{1, 2, 3, 5, 99, 0},
			expected: []int{1, 2, 3, 5, 99, 8},
		},
		{
			name:     "add direct indirect",
			code:     []int{101, 3, 3, 5, 99, 0},
			expected: []int{101, 3, 3, 5, 99, 8},
		},
		{
			name:     "add indirect direct",
			code:     []int{1001, 2, 3, 5, 99, 0},
			expected: []int{1001, 2, 3, 5, 99, 6},
		},
		{
			name:     "add direct direct",
			code:     []int{1101, 2, 3, 5, 99, 0},
			expected: []int{1101, 2, 3, 5, 99, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intcode := NewIntCode(tt.code, os.Stdin, os.Stdout)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, intcode.Code())
		})
	}
}

func TestRunOpMult(t *testing.T) {
	tests := []struct {
		name     string
		code     []int
		expected []int
	}{
		{
			name:     "mult indirect indirect",
			code:     []int{2, 2, 3, 5, 99, 0},
			expected: []int{2, 2, 3, 5, 99, 15},
		},
		{
			name:     "mult direct indirect",
			code:     []int{102, 3, 3, 5, 99, 0},
			expected: []int{102, 3, 3, 5, 99, 15},
		},
		{
			name:     "mult indirect direct",
			code:     []int{1002, 2, 3, 5, 99, 0},
			expected: []int{1002, 2, 3, 5, 99, 9},
		},
		{
			name:     "mult direct direct",
			code:     []int{1102, 2, 3, 5, 99, 0},
			expected: []int{1102, 2, 3, 5, 99, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intcode := NewIntCode(tt.code, os.Stdin, os.Stdout)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, intcode.Code())
		})
	}
}

func TestRunOpAddIndirectIndirectMultIndirectIndirect(t *testing.T) {
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
	tests := []struct {
		name     string
		code     []int
		expected []int
	}{
		{
			name:     "output indirect",
			code:     []int{4, 3, 99, 13},
			expected: []int{4, 3, 99, 13},
		},
		{
			name:     "output direct",
			code:     []int{104, 13, 99, 0},
			expected: []int{104, 13, 99, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer

			intcode := NewIntCode(tt.code, os.Stdin, &stdout)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, intcode.Code())
			assert.Equal(t, []byte("13\n"), stdout.Bytes())
		})
	}
}

func TestRunMisc(t *testing.T) {
	tests := []struct {
		name     string
		code     []int
		expected []int
	}{
		{
			name:     "negative immediate value",
			code:     []int{1101, 100, -1, 4, 0},
			expected: []int{1101, 100, -1, 4, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer

			intcode := NewIntCode(tt.code, os.Stdin, &stdout)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, intcode.Code())
		})
	}
}
