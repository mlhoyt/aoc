package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// IntCode represents stored program code with a program counter
type IntCode struct {
	code      []int
	inputSrc  io.Reader
	outputSrc io.Writer
	pc        int
}

// NewIntCode takes a stored program code and returns an IntCode object
func NewIntCode(intcode []int, inputSrc io.Reader, outputSrc io.Writer) *IntCode {
	newIntCode := IntCode{
		code:      intcode,
		inputSrc:  inputSrc,
		outputSrc: outputSrc,
		pc:        0,
	}

	return &newIntCode
}

// Run repeatedly step through the stored program until an error or the termination code is reached
func (u *IntCode) Run() error {
	var err error

	for running := true; running; {
		running, err = u.step()
		if err != nil {
			return err
		}
	}

	return nil
}

type instruction struct {
	kind         IntCodeInstructionKind
	argAddrModes []intCodeAddressingMode
}

func (u instruction) Kind() IntCodeInstructionKind {
	return u.kind
}

func (u instruction) AddressingModes() []intCodeAddressingMode {
	return u.argAddrModes
}

func newInstruction(op int) instruction {
	kind := IntCodeInstructionKind(op % 100)
	op = op / 100

	argAddrModes := []intCodeAddressingMode{}
	for op > 0 {
		argAddrModes = append(argAddrModes, intCodeAddressingMode(op%10))
		op = op / 10
	}

	return instruction{
		kind:         kind,
		argAddrModes: argAddrModes,
	}
}

type IntCodeInstructionKind int

const (
	IntCodeInstructionKindAdd    IntCodeInstructionKind = 1
	IntCodeInstructionKindMul    IntCodeInstructionKind = 2
	IntCodeInstructionKindInput  IntCodeInstructionKind = 3
	IntCodeInstructionKindOutput IntCodeInstructionKind = 4
	IntCodeInstructionKindHalt   IntCodeInstructionKind = 99
)

type intCodeAddressingMode int

const (
	intCodeAddressingModeUnknown   intCodeAddressingMode = -1
	intCodeAddressingModeIndirect  intCodeAddressingMode = 0
	intCodeAddressingModeImmediate intCodeAddressingMode = 1
)

// Step attempts to execute the next instruction in the stored program (at the current program counter)
func (u *IntCode) step() (bool, error) {
	instr := u.nextInstruction()

	switch instr.Kind() {
	case IntCodeInstructionKindAdd:
		cmd := func(x int, y int) int { return x + y }
		if err := u.executeBinaryOperation(cmd, instr.AddressingModes()); err != nil {
			return false, err
		}
		u.pc += 4
		return true, nil
	case IntCodeInstructionKindMul:
		cmd := func(x int, y int) int { return x * y }
		if err := u.executeBinaryOperation(cmd, instr.AddressingModes()); err != nil {
			return false, err
		}
		u.pc += 4
		return true, nil
	case IntCodeInstructionKindInput:
		if err := u.executeInputOperation(); err != nil {
			return false, err
		}
		u.pc += 2
		return true, nil
	case IntCodeInstructionKindOutput:
		if err := u.executeOutputOperation(instr.AddressingModes()); err != nil {
			return false, err
		}
		u.pc += 2
		return true, nil
	case IntCodeInstructionKindHalt:
		u.pc++
		return false, nil
	default:
		return false, fmt.Errorf("unexpected opcode: %d", u.code[u.pc])
	}
}

func (u *IntCode) nextInstruction() instruction {
	return newInstruction(u.code[u.pc])
}

func (u *IntCode) executeBinaryOperation(cmd func(int, int) int, argAddrModes []intCodeAddressingMode) error {
	safeArgAddrModes := make([]intCodeAddressingMode, 3)
	copy(safeArgAddrModes, argAddrModes)
	safeArgAddrModes[2] = intCodeAddressingModeImmediate

	args := []int{}
	pc := u.pc + 1
	for _, mode := range safeArgAddrModes {
		arg, err := u.getMemAt(pc, mode)
		if err != nil {
			return err
		}

		pc++

		args = append(args, arg)
	}

	return u.setMemAt(args[2], cmd(args[0], args[1]))
}

func (u *IntCode) executeInputOperation() error {
	safeArgAddrModes := make([]intCodeAddressingMode, 1)
	safeArgAddrModes[0] = intCodeAddressingModeImmediate

	args := []int{}
	pc := u.pc + 1
	for _, mode := range safeArgAddrModes {
		arg, err := u.getMemAt(pc, mode)
		if err != nil {
			return err
		}

		pc++

		args = append(args, arg)
	}

	reader := bufio.NewReader(u.inputSrc)
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	val, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return err
	}

	// valBytes := []byte{}
	// _, err := u.inputSrc.Read(valBytes)
	// if err != nil {
	// 	return err
	// }
	// valStr := strings.TrimSpace(string(valBytes))

	// val, err := strconv.Atoi(valStr)
	// if err != nil {
	// 	return err
	// }

	return u.setMemAt(args[0], val)
}

func (u *IntCode) executeOutputOperation(argAddrModes []intCodeAddressingMode) error {
	safeArgAddrModes := make([]intCodeAddressingMode, 1)
	copy(safeArgAddrModes, argAddrModes)

	args := []int{}
	pc := u.pc + 1
	for _, mode := range safeArgAddrModes {
		arg, err := u.getMemAt(pc, mode)
		if err != nil {
			return err
		}

		pc++

		args = append(args, arg)
	}

	valStr := fmt.Sprintf("%d\n", args[0])
	_, err := u.outputSrc.Write([]byte(valStr))
	return err
}

// getMemAt returns the value in program code at given address with the given address mode
func (u *IntCode) getMemAt(n int, mode intCodeAddressingMode) (int, error) {
	switch mode {
	case intCodeAddressingModeIndirect:
		v := u.code[u.code[n]]
		return v, nil
	case intCodeAddressingModeImmediate:
		v := u.code[n]
		return v, nil
	default:
		return 0, fmt.Errorf("")
	}
}

// setMemAt sets the value in program code at given address with the given value
func (u *IntCode) setMemAt(n int, v int) error {
	if n >= len(u.code) {
		return fmt.Errorf("address %d out of range", n)
	}

	u.code[n] = v

	return nil
}

// MemAt returns the current value in program code at given address
func (u *IntCode) MemAt(n int) int {
	if n >= len(u.code) {
		return 0
	}

	return u.code[n]
}

// Code returns the current stored program code
func (u *IntCode) Code() []int {
	return u.code
}
