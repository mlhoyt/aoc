package intcode

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

// Run repeatedly step through the stored program until an error or the termination code is reached
func (u *IntCode) Run() error {
	for {
		running, err := u.step()
		if err != nil {
			return err
		}

		if !running {
			break
		}
	}

	return nil
}

type IntCodeInstr int

const (
	IntCodeInstrAdd  IntCodeInstr = 1
	IntCodeInstrMul  IntCodeInstr = 2
	IntCodeInstrHalt IntCodeInstr = 99
)

// Step attempts to execute the next instruction in the stored program (at the current program counter)
func (u *IntCode) step() (bool, error) {
	switch IntCodeInstr(u.code[u.pc]) {
	case IntCodeInstrAdd:
		// u.SetMemAt(u.MemAt(u.pc+3), (u.MemAt(u.MemAt(u.pc+1)) + u.MemAt(u.MemAt(u.pc+2))))
		u.code[u.code[u.pc+3]] = u.code[u.code[u.pc+1]] + u.code[u.code[u.pc+2]]
		u.pc += 4
		return true, nil
	case IntCodeInstrMul:
		u.code[u.code[u.pc+3]] = u.code[u.code[u.pc+1]] * u.code[u.code[u.pc+2]]
		u.pc += 4
		return true, nil
	case IntCodeInstrHalt:
		u.pc++
		return false, nil
	default:
		return false, fmt.Errorf("unexpected opcode: %d", u.code[u.pc])
	}
}

// MemAt returns the current value in program code at given address
func (u *IntCode) MemAt(n int) int {
	if n >= len(u.code) {
		return 0
	}

	return u.code[n]
}

// SetMemAt sets the current value in program code at given address with the given value
func (u *IntCode) SetMemAt(n int, v int) error {
	if n >= len(u.code) {
		return fmt.Errorf("address %d out of range", n)
	}

	u.code[n] = v

	return nil
}

// Code returns the current stored program code
func (u *IntCode) Code() []int {
	return u.code
}
