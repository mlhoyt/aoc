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

// Code returns the current stored program code
func (u *IntCode) Code() []int {
	return u.code
}

// Step attempts to execute the next instruction in the stored program (at the current program counter)
func (u *IntCode) Step() (bool, error) {
	// fmt.Printf("[DEBUG] IntCode::Step: code=%v pc=%d\n", u.code, u.pc)
	// fmt.Printf("[DEBUG] IntCode::Step: opcode=%d\n", u.code[u.pc])

	switch u.code[u.pc] {
	case 1:
		u.code[u.code[u.pc+3]] = u.code[u.code[u.pc+1]] + u.code[u.code[u.pc+2]]
		u.pc += 4
		return true, nil
	case 2:
		u.code[u.code[u.pc+3]] = u.code[u.code[u.pc+1]] * u.code[u.code[u.pc+2]]
		u.pc += 4
		return true, nil
	case 99:
		u.pc++
		return false, nil
	default:
		return false, fmt.Errorf("unexpected opcode: %d", u.code[u.pc])
	}
}

// Run attempts to repeatedly step through the stored program until the termination code is reached
func (u *IntCode) Run() error {
	for {
		running, err := u.Step()
		if err != nil {
			return err
		}

		if !running {
			break
		}
	}

	return nil
}
