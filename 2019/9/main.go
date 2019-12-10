package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

var debug = false

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(data)), ",")

	debug = false
	part1(input)
	debug = false
	part2(input)
}

func part1(input []string) {
	is := toInt(input)
	is = append(is, make([]int64, 100000)...)
	in := make(chan int64, 1)
	output := make(chan int64, 100)
	in <- 1
	run(0, is, in, output, nil)
	s := len(output)
	for i := 0; i < s; i++ {
		fmt.Print(<-output)
		fmt.Print(" ")
	}
	fmt.Println()
}

func part2(input []string) {
	is := toInt(input)
	is = append(is, make([]int64, 100000)...)
	in := make(chan int64, 1)
	output := make(chan int64, 100)
	in <- 2
	run(0, is, in, output, nil)
	s := len(output)
	for i := 0; i < s; i++ {
		fmt.Print(<-output)
		fmt.Print(" ")
	}
	fmt.Println()
}

type opCode int64

const (
	opAdd  opCode = 1
	opMul  opCode = 2
	opIn   opCode = 3
	opOut  opCode = 4
	opJiT  opCode = 5
	opJiF  opCode = 6
	opLess opCode = 7
	opEq   opCode = 8
	opRel  opCode = 9
	opHalt opCode = 99
)

func toInt(lines []string) []int64 {
	ret := []int64{}

	for _, l := range lines {
		i, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}

	return ret
}

func opStr(o opCode) string {
	switch o {
	case opAdd:
		return "add"
	case opMul:
		return "mul"
	case opIn:
		return "in"
	case opOut:
		return "out"
	case opJiT:
		return "jit"
	case opJiF:
		return "jif"
	case opLess:
		return "less"
	case opEq:
		return "eq"
	case opRel:
		return "rb"
	case opHalt:
		return "halt"
	default:
		return "%"
	}
}

func println(iface ...interface{}) {
	if debug {
		fmt.Println(iface...)
	}
}

func print(iface ...interface{}) {
	if debug {
		fmt.Print(iface...)
	}
}

func get(is []int64, mem map[int64]int64, i int64) int64 {
	val := int64(0)
	if int(i) < len(is) {
		val = is[i]
	} else {
		val = mem[i]
	}
	println("LOAD &"+strconv.FormatInt(i, 10), ">", val)
	return val
}

func set(is []int64, mem map[int64]int64, i, v int64) {
	if int(i) < len(is) {
		is[i] = v
	} else {
		mem[i] = v
	}
	println("STORE &"+strconv.FormatInt(i, 10), "<", v)
}

func run(id int64, is []int64, input, output chan int64, wg *sync.WaitGroup) {
	pc := int64(0)
	base := int64(0)
	mem := map[int64]int64{}
	for {
		println()
		op, modes := decodeOp(get(is, mem, pc))
		println("op", opStr(op), modes)
		switch op {
		case opAdd:
			v1 := value(is, mem, modes[2], get(is, mem, pc+1), base)
			v2 := value(is, mem, modes[1], get(is, mem, pc+2), base)
			v3 := valueForStore(is, mem, modes[0], get(is, mem, pc+3), base)
			println("add", v1, v2, "set", v3)
			set(is, mem, v3, v1+v2)
			pc += 4
		case opMul:
			v1 := value(is, mem, modes[2], get(is, mem, pc+1), base)
			v2 := value(is, mem, modes[1], get(is, mem, pc+2), base)
			v3 := valueForStore(is, mem, modes[0], get(is, mem, pc+3), base)
			println("mul", v1, v2, "set", v3)
			set(is, mem, v3, v1*v2)
			pc += 4
		case opIn:
			i := valueForStore(is, mem, modes[0], get(is, mem, pc+1), base)
			v := <-input
			println("input set", i, v)
			set(is, mem, i, v)
			pc += 2
		case opOut:
			v := value(is, mem, modes[0], get(is, mem, pc+1), base)
			println("output", v, pc+1)
			output <- v
			pc += 2
		case opJiT:
			v1 := value(is, mem, modes[1], get(is, mem, pc+1), base)
			v2 := value(is, mem, modes[0], get(is, mem, pc+2), base)
			println("jit", v1, v2)
			if v1 != 0 {
				pc = v2
			} else {
				pc += 3
			}
		case opJiF:
			v1 := value(is, mem, modes[1], get(is, mem, pc+1), base)
			v2 := value(is, mem, modes[0], get(is, mem, pc+2), base)
			println("jif", v1, v2)
			if v1 == 0 {
				pc = v2
			} else {
				pc += 3
			}
		case opLess:
			v1 := value(is, mem, modes[2], get(is, mem, pc+1), base)
			v2 := value(is, mem, modes[1], get(is, mem, pc+2), base)
			v3 := valueForStore(is, mem, modes[0], get(is, mem, pc+3), base)
			println("less", v1, v2, "set", v3)
			if v1 < v2 {
				set(is, mem, v3, 1)
			} else {
				set(is, mem, v3, 0)
			}
			pc += 4
		case opEq:
			v1 := value(is, mem, modes[2], get(is, mem, pc+1), base)
			v2 := value(is, mem, modes[1], get(is, mem, pc+2), base)
			v3 := valueForStore(is, mem, modes[0], get(is, mem, pc+3), base)
			println("eq", v1, v2, "set", v3)
			if v1 == v2 {
				set(is, mem, v3, 1)
			} else {
				set(is, mem, v3, 0)
			}
			pc += 4
		case opRel:
			v := value(is, mem, modes[0], get(is, mem, pc+1), base)
			println("base", base, v, base+v)
			base += v
			pc += 2
		case opHalt:
			if wg != nil {
				wg.Done()
			}
			return
		default:
			panic("unsupported op")
		}
	}
}

func value(is []int64, mem map[int64]int64, mode, arg, base int64) int64 {
	if mode == 0 {
		return get(is, mem, arg)
	} else if mode == 1 {
		return arg
	} else if mode == 2 {
		return get(is, mem, base+arg)
	}
	panic("Unsupported mode")
}

func valueForStore(is []int64, mem map[int64]int64, mode, arg, base int64) int64 {
	if mode == 0 {
		return arg
	} else if mode == 2 {
		return arg + base
	}
	panic("unsupported mode")
}

func decodeOp(i int64) (opCode, []int64) {
	chars := strconv.FormatInt(i, 10)
	digits := []int64{}
	for _, c := range chars {
		d, _ := strconv.ParseInt(string(c), 10, 64)
		digits = append(digits, d)
	}
	if len(digits) == 1 {
		digits = append([]int64{0}, digits...)
	}

	op := opCode(digits[len(digits)-2]*10 + digits[len(digits)-1])
	pad := 0
	switch op {
	case opAdd, opMul, opLess, opEq:
		pad = 5
	case opJiT, opJiF:
		pad = 4
	case opIn, opOut, opRel:
		pad = 3
	case opHalt:
		pad = 0
	default:
		fmt.Println(i, op)
		panic("unknown op")
	}

	for i := len(digits); i < pad; i++ {
		digits = append([]int64{0}, digits...)
	}
	return op, digits[:len(digits)-2]
}
