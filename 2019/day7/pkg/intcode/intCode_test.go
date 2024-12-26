package intcode

import (
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
			input := make(IOSrc)
			output := make(IOSrc)

			intcode := NewIntCode(tt.code, input, output)
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
			input := make(IOSrc)
			output := make(IOSrc)

			intcode := NewIntCode(tt.code, input, output)
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

	input := make(IOSrc)
	output := make(IOSrc)

	intcode := NewIntCode(code, input, output)
	err := intcode.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, intcode.Code())
}

func TestRunOpInput(t *testing.T) {
	tests := []struct {
		name     string
		code     []int
		input    []int
		expected []int
	}{
		{
			name:     "single read indirect",
			code:     []int{3, 1, 99},
			input:    []int{13},
			expected: []int{3, 13, 99},
		},
		{
			name:     "multiple read indirect",
			code:     []int{3, 1, 3, 3, 99},
			input:    []int{13, 31},
			expected: []int{3, 13, 3, 31, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(IOSrc, 10)
			output := make(IOSrc, 10)

			for _, v := range tt.input {
				input <- v
			}
			intcode := NewIntCode(tt.code, input, output)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, intcode.Code())
		})
	}
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
			input := make(IOSrc, 10)
			output := make(IOSrc, 10)

			intcode := NewIntCode(tt.code, input, output)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.expected, intcode.Code())
			assert.Equal(t, 13, <-output)
		})
	}
}

func TestRunMisc(t *testing.T) {
	tests := []struct {
		name          string
		code          []int
		expected      []int
		checkExpected bool
		input         []int
		output        []int
	}{
		{
			name:          "negative immediate value",
			code:          []int{1101, 100, -1, 4, 0},
			expected:      []int{1101, 100, -1, 4, 99},
			checkExpected: true,
			input:         []int{},
			output:        []int{},
		},
		{
			name:          "is input equal to 8, indirect, true",
			code:          []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			expected:      []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8},
			checkExpected: true,
			input:         []int{8},
			output:        []int{1},
		},
		{
			name:          "is input equal to 8, indirect, false",
			code:          []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			expected:      []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8},
			checkExpected: true,
			input:         []int{7},
			output:        []int{0},
		},
		{
			name:          "is input equal to 8, immediate, true",
			code:          []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			expected:      []int{3, 3, 1108, 1, 8, 3, 4, 3, 99},
			checkExpected: true,
			input:         []int{8},
			output:        []int{1},
		},
		{
			name:          "is input equal to 8, immediate, false",
			code:          []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			expected:      []int{3, 3, 1108, 0, 8, 3, 4, 3, 99},
			checkExpected: true,
			input:         []int{7},
			output:        []int{0},
		},
		{
			name:          "is input less than 8, indirect, true",
			code:          []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			expected:      []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8},
			checkExpected: true,
			input:         []int{7},
			output:        []int{1},
		},
		{
			name:          "is input less than 8, indirect, false",
			code:          []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			expected:      []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8},
			checkExpected: true,
			input:         []int{8},
			output:        []int{0},
		},
		{
			name:          "is input less than 8, immediate, true",
			code:          []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			expected:      []int{3, 3, 1107, 1, 8, 3, 4, 3, 99},
			checkExpected: true,
			input:         []int{7},
			output:        []int{1},
		},
		{
			name:          "is input less than 8, immediate, false",
			code:          []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			expected:      []int{3, 3, 1107, 0, 8, 3, 4, 3, 99},
			checkExpected: true,
			input:         []int{8},
			output:        []int{0},
		},
		{
			name:          "is input equal to 0, indirect, true",
			code:          []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			expected:      []int{},
			checkExpected: false,
			input:         []int{0},
			output:        []int{0},
		},
		{
			name:          "is input equal to 0, indirect, false",
			code:          []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			expected:      []int{},
			checkExpected: false,
			input:         []int{1},
			output:        []int{1},
		},
		{
			name:          "is input equal to 0, immediate, true",
			code:          []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			expected:      []int{},
			checkExpected: false,
			input:         []int{0},
			output:        []int{0},
		},
		{
			name:          "is input equal to 0, immediate, false",
			code:          []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			expected:      []int{},
			checkExpected: false,
			input:         []int{1},
			output:        []int{1},
		},
		{
			name:          "input relative to 8, below",
			code:          []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			expected:      []int{},
			checkExpected: false,
			input:         []int{7},
			output:        []int{999},
		},
		{
			name:          "input relative to 8, equal",
			code:          []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			expected:      []int{},
			checkExpected: false,
			input:         []int{8},
			output:        []int{1000},
		},
		{
			name:          "input relative to 8, above",
			code:          []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			expected:      []int{},
			checkExpected: false,
			input:         []int{9},
			output:        []int{1001},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(IOSrc, 10)
			output := make(IOSrc, 10)

			for _, v := range tt.input {
				input <- v
			}
			intcode := NewIntCode(tt.code, input, output)
			err := intcode.Run()
			if err != nil {
				t.Error(err)
			}

			if tt.checkExpected {
				assert.Equal(t, tt.expected, intcode.Code())
			}

			for _, v := range tt.output {
				assert.Equal(t, v, <-output)
			}
		})
	}
}
