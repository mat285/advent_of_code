package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), ",")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	is := toInt(lines)
	fmt.Println(run(is, 1, false))
}

func part2(lines []string) {
	is := toInt(lines)
	fmt.Println(run(is, 5, false))
}

type opCode int

const (
	opAdd  opCode = 1
	opMul  opCode = 2
	opIn   opCode = 3
	opOut  opCode = 4
	opJiT  opCode = 5
	opJiF  opCode = 6
	opLess opCode = 7
	opEq   opCode = 8
	opHalt opCode = 99
)

func toInt(lines []string) []int {
	ret := []int{}

	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}

	return ret
}

func run(is []int, in int, debug bool) int {
	pc := 0
	out := -1
	for {
		op, modes := decodeOp(is[pc])
		switch op {
		case opAdd:
			v1 := value(is, modes[2], is[pc+1])
			v2 := value(is, modes[1], is[pc+2])
			v3 := is[pc+3]
			is[v3] = v1 + v2
			pc += 4
		case opMul:
			v1 := value(is, modes[2], is[pc+1])
			v2 := value(is, modes[1], is[pc+2])
			v3 := is[pc+3]
			is[v3] = v1 * v2
			pc += 4
		case opIn:
			v := is[pc+1]
			is[v] = in
			pc += 2
		case opOut:
			v := value(is, modes[0], is[pc+1])
			if debug {
				fmt.Println(v)
			}
			out = v
			pc += 2
		case opJiT:
			v1 := value(is, modes[1], is[pc+1])
			v2 := value(is, modes[0], is[pc+2])
			if v1 != 0 {
				pc = v2
			} else {
				pc += 3
			}
		case opJiF:
			v1 := value(is, modes[1], is[pc+1])
			v2 := value(is, modes[0], is[pc+2])
			if v1 == 0 {
				pc = v2
			} else {
				pc += 3
			}
		case opLess:
			v1 := value(is, modes[2], is[pc+1])
			v2 := value(is, modes[1], is[pc+2])
			v3 := is[pc+3]
			if v1 < v2 {
				is[v3] = 1
			} else {
				is[v3] = 0
			}
			pc += 4
		case opEq:
			v1 := value(is, modes[2], is[pc+1])
			v2 := value(is, modes[1], is[pc+2])
			v3 := is[pc+3]
			if v1 == v2 {
				is[v3] = 1
			} else {
				is[v3] = 0
			}
			pc += 4
		case opHalt:
			return out
		default:
			panic("unsupported op")
		}
	}
}

func value(is []int, mode, arg int) int {
	if mode == 0 {
		return is[arg]
	} else if mode == 1 {
		return arg
	}
	panic("Unsupported mode")
}

func decodeOp(i int) (opCode, []int) {
	chars := strconv.Itoa(i)
	digits := []int{}
	for _, c := range chars {
		d, _ := strconv.Atoi(string(c))
		digits = append(digits, d)
	}
	if len(digits) == 1 {
		digits = append([]int{0}, digits...)
	}
	op := opCode(digits[len(digits)-2]*10 + digits[len(digits)-1])
	pad := 0
	switch op {
	case opAdd, opMul, opLess, opEq:
		pad = 5
	case opJiT, opJiF:
		pad = 4
	case opIn, opOut:
		pad = 3
	case opHalt:
		pad = 0
	default:
		panic("unknown op")
	}

	for i := len(digits); i < pad; i++ {
		digits = append([]int{0}, digits...)
	}
	return op, digits[:len(digits)-2]
}
