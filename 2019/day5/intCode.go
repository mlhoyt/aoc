package main

import (
	"fmt"
)

// IntCode represents stored program code with a program counter
type IntCode struct {
	code []int
	pc   int
}

// NewIntCode takes a stored program code and returns an IntCode object
func NewIntCode(intcode []int) *IntCode {
	newIntCode := IntCode{
		code: intcode,
		pc:   0,
	}

	return &newIntCode
}

// Run repeatedly steps through the stored program until the termination code is reached (or an error is encountered)
func (u *IntCode) Run() error {
	var err error

	cont := true
	for cont {
		cont, err = u.step()
		if err != nil {
			return err
		}
	}

	return nil
}

// Step executes the next instruction
func (u *IntCode) step() (bool, error) {
	switch u.op() {
	case opcodeAddIndirect:
		return u.opAddIndirect(), nil
	case opcodeMultIndirect:
		return u.opMultIndirect(), nil
	case opcodeInput:
		return u.opInput(), nil
	case opcodeOutput:
		return u.opOutput(), nil
	case opcodeEnd:
		return u.opEnd(), nil
	default:
		return false, fmt.Errorf("unexpected opcode: %d", u.code[u.pc])
	}
}

func (u *IntCode) op() opcode {
	return NewOpcode(u.code[u.pc])
}

func (u *IntCode) pcIncr() {
	u.pc++
}

func (u *IntCode) getIndirect() int {
	v := u.code[u.code[u.pc]]
	u.pcIncr()

	return v
}

func (u *IntCode) putIndirect(v int) {
	u.code[u.code[u.pc]] = v
	u.pcIncr()
}

func (u *IntCode) opAddIndirect() bool {
	u.pcIncr()
	arg1 := u.getIndirect()
	arg2 := u.getIndirect()
	u.putIndirect(arg1 + arg2)

	return true
}

func (u *IntCode) opMultIndirect() bool {
	u.pcIncr()
	arg1 := u.getIndirect()
	arg2 := u.getIndirect()
	u.putIndirect(arg1 * arg2)

	return true
}

func (u *IntCode) opInput() bool {
	u.pcIncr()
	v := 13 // FIXME: need to read from stdin
	u.putIndirect(v)

	return true
}

func (u *IntCode) opOutput() bool {
	u.pcIncr()
	v := u.getIndirect()
	fmt.Println(v)

	return true
}

func (u *IntCode) opEnd() bool {
	u.pcIncr()

	return false
}

// Code returns the current stored program code
func (u *IntCode) Code() []int {
	return u.code
}

type opcode int

const (
	opcodeUnknown      opcode = 0
	opcodeAddIndirect         = 1
	opcodeMultIndirect        = 2
	opcodeInput               = 3
	opcodeOutput              = 4
	opcodeEnd                 = 99
)

var opcodeEnumMap = map[int]opcode{
	1:  opcodeAddIndirect,
	2:  opcodeMultIndirect,
	3:  opcodeInput,
	4:  opcodeOutput,
	99: opcodeEnd,
}

func NewOpcode(n int) opcode {
	if v, ok := opcodeEnumMap[n]; ok {
		return v
	}

	return opcodeUnknown
}

// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
