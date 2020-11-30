package intcode

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunAddIndirect(t *testing.T) {
	code := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}

	computer := NewIntCode(code, nil, nil)
	err := computer.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, computer.Code())
}

func TestRuMulIndirect(t *testing.T) {
	code := []int{2, 3, 0, 3, 99}
	expected := []int{2, 3, 0, 6, 99}

	computer := NewIntCode(code, nil, nil)
	err := computer.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, computer.Code())
}

func TestRunMulIndirect2(t *testing.T) {
	code := []int{2, 4, 4, 5, 99, 0}
	expected := []int{2, 4, 4, 5, 99, 9801}

	computer := NewIntCode(code, nil, nil)
	err := computer.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, computer.Code())
}

func TestRunAddIndirectMulIndirect(t *testing.T) {
	code := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	computer := NewIntCode(code, nil, nil)
	err := computer.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, computer.Code())
}

func TestRunAddIndirectMulIndirect2(t *testing.T) {
	code := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	expected := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}

	computer := NewIntCode(code, nil, nil)
	err := computer.Run()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, computer.Code())
}

func TestNewInstruction(t *testing.T) {
	testcases := []struct {
		name            string
		code            int
		kind            IntCodeInstructionKind
		addressingModes []intCodeAddressingMode
	}{
		{
			name:            "ADD INDIRECT, INDIRECT",
			code:            1,
			kind:            IntCodeInstructionKindAdd,
			addressingModes: []intCodeAddressingMode{},
		},
		{
			name: "ADD IMMEDIATE, INDIRECT",
			code: 101,
			kind: IntCodeInstructionKindAdd,
			addressingModes: []intCodeAddressingMode{
				intCodeAddressingModeImmediate,
			},
		},
		{
			name: "ADD IMMEDIATE, IMMEDIATE",
			code: 1101,
			kind: IntCodeInstructionKindAdd,
			addressingModes: []intCodeAddressingMode{
				intCodeAddressingModeImmediate,
				intCodeAddressingModeImmediate,
			},
		},
		{
			name:            "MUL INDIRECT, INDIRECT",
			code:            2,
			kind:            IntCodeInstructionKindMul,
			addressingModes: []intCodeAddressingMode{},
		},
		{
			name: "MUL IMMEDIATE, INDIRECT",
			code: 102,
			kind: IntCodeInstructionKindMul,
			addressingModes: []intCodeAddressingMode{
				intCodeAddressingModeImmediate,
			},
		},
		{
			name: "MUL IMMEDIATE, IMMEDIATE",
			code: 1102,
			kind: IntCodeInstructionKindMul,
			addressingModes: []intCodeAddressingMode{
				intCodeAddressingModeImmediate,
				intCodeAddressingModeImmediate,
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			actual := newInstruction(testcase.code)

			assert.Equal(t, testcase.kind, actual.Kind())
			assert.Equal(t, testcase.addressingModes, actual.AddressingModes())
		})
	}
}

func TestGetMemAt(t *testing.T) {
	testcases := []struct {
		name        string
		code        []int
		loc         int
		addressMode intCodeAddressingMode
		expected    int
	}{
		{
			name:        "indirect 0 -> 1 -> 2",
			code:        []int{1, 2, 4, 8, 16},
			loc:         0,
			addressMode: intCodeAddressingModeIndirect,
			expected:    2,
		},
		{
			name:        "indirect 1 -> 2 -> 4",
			code:        []int{1, 2, 4, 8, 16},
			loc:         1,
			addressMode: intCodeAddressingModeIndirect,
			expected:    4,
		},
		{
			name:        "immediate 2 -> 4",
			code:        []int{1, 2, 4, 8, 16},
			loc:         2,
			addressMode: intCodeAddressingModeImmediate,
			expected:    4,
		},
		{
			name:        "indirect 3 -> 8",
			code:        []int{1, 2, 4, 8, 16},
			loc:         3,
			addressMode: intCodeAddressingModeImmediate,
			expected:    8,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			computer := NewIntCode(testcase.code, nil, nil)
			actual, err := computer.getMemAt(testcase.loc, testcase.addressMode)

			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func TestSetMemAt(t *testing.T) {
	testcases := []struct {
		name     string
		code     []int
		loc      int
		value    int
		expected []int
	}{
		{
			name:     "set 0 -> 13",
			code:     []int{1, 2, 4, 8, 16},
			loc:      0,
			value:    13,
			expected: []int{13, 2, 4, 8, 16},
		},
		{
			name:     "set 1 -> 113",
			code:     []int{1, 2, 4, 8, 16},
			loc:      1,
			value:    113,
			expected: []int{1, 113, 4, 8, 16},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			computer := NewIntCode(testcase.code, nil, nil)
			err := computer.setMemAt(testcase.loc, testcase.value)

			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, computer.Code())
		})
	}
}

func TestExecuteBinaryOperation(t *testing.T) {
	testcases := []struct {
		name            string
		code            []int
		cmd             func(int, int) int
		addressingModes []intCodeAddressingMode
		expected        []int
	}{
		{
			name:            "add 0 -> 1, 1 -> 0, 5",
			code:            []int{1, 0, 1, 5, 99, 0},
			cmd:             func(x int, y int) int { return x + y },
			addressingModes: []intCodeAddressingMode{},
			expected:        []int{1, 0, 1, 5, 99, 1},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			computer := NewIntCode(testcase.code, nil, nil)
			err := computer.executeBinaryOperation(testcase.cmd, testcase.addressingModes)

			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, computer.Code())
		})
	}
}

func TestExecuteInputOperation(t *testing.T) {
	testcases := []struct {
		name     string
		code     []int
		input    string
		expected []int
	}{
		{
			name:     "input 13 -> indirect",
			code:     []int{3, 3, 99, 0},
			input:    "13\n",
			expected: []int{3, 3, 99, 13},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			inputSrc := bytes.NewBufferString(testcase.input)
			computer := NewIntCode(testcase.code, inputSrc, nil)
			err := computer.executeInputOperation()

			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, computer.Code())
		})
	}
}

func TestExecuteOutputOperation(t *testing.T) {
	testcases := []struct {
		name            string
		code            []int
		addressingModes []intCodeAddressingMode
		expected        string
	}{
		{
			name:            "output indirect -> 13",
			code:            []int{4, 3, 99, 13},
			addressingModes: []intCodeAddressingMode{},
			expected:        "13\n",
		},
		{
			name:            "output immedaite -> 13",
			code:            []int{104, 13, 99},
			addressingModes: []intCodeAddressingMode{intCodeAddressingModeImmediate},
			expected:        "13\n",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			outputSrc := bytes.Buffer{}
			computer := NewIntCode(testcase.code, nil, &outputSrc)
			err := computer.executeOutputOperation(testcase.addressingModes)

			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, outputSrc.String())
		})
	}
}
