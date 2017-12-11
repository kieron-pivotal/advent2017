package registers

import (
	"regexp"
	"strconv"
)

type operation int

const (
	DEC operation = iota
	INC
)

type comparison int

const (
	GT comparison = iota
	GTE
	LT
	LTE
	EQ
	NEQ
)

type Set struct {
	Registers    map[string]int
	Instructions []Instruction
}

type Instruction struct {
	Register   string
	Op         operation
	Amt        int
	CondTarget string
	Comparison comparison
	CompAmt    int
}

func (s Set) AddInstruction(line string) {
	instr := ParseInstruction(line)
	s.Instructions = append(s.Instructions, instr)
	s.Registers[instr.Register] = 0
}

func ParseInstruction(line string) Instruction {
	re := regexp.MustCompile("([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) ([<>=!]+) (-?[0-9]+)")
	parts := re.FindStringSubmatch(line)
	instr := Instruction{
		Register:   parts[1],
		CondTarget: parts[4],
	}

	if parts[2] == "inc" {
		instr.Op = INC
	}

	amt, _ := strconv.Atoi(parts[3])
	instr.Amt = amt

	amt, _ = strconv.Atoi(parts[6])
	instr.CompAmt = amt

	switch parts[5] {
	case ">":
		instr.Comparison = GT
	case "<":
		instr.Comparison = LT
	case ">=":
		instr.Comparison = GTE
	case "<=":
		instr.Comparison = LTE
	case "==":
		instr.Comparison = EQ
	case "!=":
		instr.Comparison = NEQ
	}
	return instr
}

func (s Set) Process() {
	for _, instr := range s.Instructions {
		cond := false
		toComp := s.Registers[instr.CondTarget]
		switch instr.Comparison {
		case GT:
			cond = toComp > instr.CompAmt
		case GTE:
			cond = toComp >= instr.CompAmt
		case LT:
			cond = toComp < instr.CompAmt
		case LTE:
			cond = toComp <= instr.CompAmt
		case EQ:
			cond = toComp == instr.CompAmt
		case NEQ:
			cond = toComp != instr.CompAmt
		}
		if cond {
			switch instr.Op {
			case INC:
				s.Registers[instr.Register] += instr.Amt
			case DEC:
				s.Registers[instr.Register] -= instr.Amt
			}
		}
	}
}

func (s Set) GetMaxRegister() int {
	max := 0
	for _, val := range s.Registers {
		if val > max {
			max = val
		}
	}
	return max
}
