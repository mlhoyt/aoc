package intcode

import (
	"fmt"
)

type IOSrc chan int

// IntCode represents stored program code with an input source, an output source, and a program counter
type IntCode struct {
	code      []int
	inputSrc  IOSrc
	outputSrc IOSrc
	pc        int
}

// NewIntCode takes a stored program code, an input source, and an output source and returns an IntCode object
func NewIntCode(code []int, inputSrc IOSrc, outputSrc IOSrc) *IntCode {
	newIntCode := IntCode{
		code:      code,
		inputSrc:  inputSrc,
		outputSrc: outputSrc,
		pc:        0,
	}

	return &newIntCode
}

// Run repeatedly steps through the stored program until the termination code is reached (or an error is encountered)
func (u *IntCode) Run() error {
	var err error

	for cont := true; cont; {
		cont, err = u.step()
		if err != nil {
			return err
		}
	}

	return nil
}

// Step executes the next instruction
func (u *IntCode) step() (bool, error) {
	instr := u.nextInstruction()

	switch instr.Op() {
	case opCodeAdd:
		return u.opAdd(instr.ArgAddrModes()), nil
	case opCodeMult:
		return u.opMult(instr.ArgAddrModes()), nil
	case opCodeInput:
		return u.opInput(instr.ArgAddrModes())
	case opCodeOutput:
		return u.opOutput(instr.ArgAddrModes())
	case opCodeJumpIfTrue:
		return u.opJumpIfTrue(instr.ArgAddrModes()), nil
	case opCodeJumpIfFalse:
		return u.opJumpIfFalse(instr.ArgAddrModes()), nil
	case opCodeLessThan:
		return u.opLessThan(instr.ArgAddrModes()), nil
	case opCodeEqual:
		return u.opEqual(instr.ArgAddrModes()), nil
	case opCodeEnd:
		return u.opEnd(), nil
	default:
		return false, fmt.Errorf("unexpected opCode: %d", u.code[u.pc])
	}
}

type instruction struct {
	op           opCode
	argAddrModes []opArgAddrMode
}

func (u *instruction) Op() opCode {
	return u.op
}

func (u *instruction) ArgAddrModes() []opArgAddrMode {
	return u.argAddrModes
}

func newInstruction(op int) instruction {
	instrOp := NewOpCode(op % 100)
	op = op / 100

	instrArgAddrModes := []opArgAddrMode{}
	for op > 0 {
		instrArgAddrModes = append(instrArgAddrModes, NewOpArgAddrMode(op%10))
		op = op / 10
	}

	return instruction{
		op:           instrOp,
		argAddrModes: instrArgAddrModes,
	}
}

func (u *IntCode) nextInstruction() instruction {
	return newInstruction(u.code[u.pc])
}

func (u *IntCode) pcIncr() {
	u.pc++
}

func (u *IntCode) pcSet(v int) {
	u.pc = v
}

func (u *IntCode) get(mode opArgAddrMode) int {
	switch mode {
	case opArgAddrModeIndirect:
		return u.getIndirect()
	case opArgAddrModeImmediate:
		return u.getImmediate()
	}

	return -1 // this really should be unreachable and error
}

func (u *IntCode) getIndirect() int {
	v := u.code[u.code[u.pc]]
	u.pcIncr()

	return v
}

func (u *IntCode) getImmediate() int {
	v := u.code[u.pc]
	u.pcIncr()

	return v
}

func (u *IntCode) putIndirect(v int) {
	u.code[u.code[u.pc]] = v
	u.pcIncr()
}

func (u *IntCode) opAdd(argAddrModes []opArgAddrMode) bool {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 3)
	u.pcIncr()

	arg1 := u.get(safeArgAddrModes[0])
	arg2 := u.get(safeArgAddrModes[1])
	u.putIndirect(arg1 + arg2)

	return true
}

func (u *IntCode) opMult(argAddrModes []opArgAddrMode) bool {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 3)
	u.pcIncr()

	arg1 := u.get(safeArgAddrModes[0])
	arg2 := u.get(safeArgAddrModes[1])
	u.putIndirect(arg1 * arg2)

	return true
}

func (u *IntCode) readFromInputSrc() (int, error) {
	input := <-u.inputSrc

	return input, nil
}

func (u *IntCode) opInput(argAddrModes []opArgAddrMode) (bool, error) {
	u.pcIncr()

	v, err := u.readFromInputSrc()
	if err != nil {
		return false, err
	}

	u.putIndirect(v)

	return true, nil
}

func (u *IntCode) writeToOutputSrc(v int) {
	u.outputSrc <- v
}

func (u *IntCode) opOutput(argAddrModes []opArgAddrMode) (bool, error) {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 1)
	u.pcIncr()

	v := u.get(safeArgAddrModes[0])

	u.writeToOutputSrc(v)

	return true, nil
}

func (u *IntCode) opJumpIfTrue(argAddrModes []opArgAddrMode) bool {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 2)
	u.pcIncr()

	arg1 := u.get(safeArgAddrModes[0])
	arg2 := u.get(safeArgAddrModes[1])

	if arg1 != 0 {
		u.pcSet(arg2)
	}

	return true
}

func (u *IntCode) opJumpIfFalse(argAddrModes []opArgAddrMode) bool {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 2)
	u.pcIncr()

	arg1 := u.get(safeArgAddrModes[0])
	arg2 := u.get(safeArgAddrModes[1])

	if arg1 == 0 {
		u.pcSet(arg2)
	}

	return true
}

func (u *IntCode) opLessThan(argAddrModes []opArgAddrMode) bool {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 3)
	u.pcIncr()

	arg1 := u.get(safeArgAddrModes[0])
	arg2 := u.get(safeArgAddrModes[1])

	v := 0
	if arg1 < arg2 {
		v = 1
	}
	u.putIndirect(v)

	return true
}

func (u *IntCode) opEqual(argAddrModes []opArgAddrMode) bool {
	safeArgAddrModes := fillArgAddrModes(argAddrModes, 3)
	u.pcIncr()

	arg1 := u.get(safeArgAddrModes[0])
	arg2 := u.get(safeArgAddrModes[1])

	v := 0
	if arg1 == arg2 {
		v = 1
	}
	u.putIndirect(v)

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

type opCode int

const (
	opCodeUnknown     opCode = 0
	opCodeAdd                = 1
	opCodeMult               = 2
	opCodeInput              = 3
	opCodeOutput             = 4
	opCodeJumpIfTrue         = 5
	opCodeJumpIfFalse        = 6
	opCodeLessThan           = 7
	opCodeEqual              = 8
	opCodeEnd                = 99
)

var opCodeEnumMap = map[int]opCode{
	1:  opCodeAdd,
	2:  opCodeMult,
	3:  opCodeInput,
	4:  opCodeOutput,
	5:  opCodeJumpIfTrue,
	6:  opCodeJumpIfFalse,
	7:  opCodeLessThan,
	8:  opCodeEqual,
	99: opCodeEnd,
}

func NewOpCode(n int) opCode {
	if v, ok := opCodeEnumMap[n]; ok {
		return v
	}

	return opCodeUnknown
}

type opArgAddrMode int

const (
	opArgAddrModeUnknown   opArgAddrMode = -1
	opArgAddrModeIndirect                = 0
	opArgAddrModeImmediate               = 1
)

var opArgAddrModeEnumMap = map[int]opArgAddrMode{
	0: opArgAddrModeIndirect,
	1: opArgAddrModeImmediate,
}

func NewOpArgAddrMode(n int) opArgAddrMode {
	if v, ok := opArgAddrModeEnumMap[n]; ok {
		return v
	}

	return opArgAddrModeUnknown
}

func fillArgAddrModes(argAddrModes []opArgAddrMode, size int) []opArgAddrMode {
	for len(argAddrModes) < size {
		argAddrModes = append(argAddrModes, opArgAddrModeIndirect)
	}

	return argAddrModes
}
