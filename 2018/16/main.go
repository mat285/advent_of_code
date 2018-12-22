package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var testLines = strings.Split(`Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`, "\n")

func main() {
	lines := read()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	tests := parseTests(lines)

	count := 0
	for _, t := range tests {
		n := possible(t)
		if len(n) >= 3 {
			count++
		}
	}
	fmt.Println(count)
}

func part2(lines []string) {
	tests := parseTests(lines)

	poss := map[int][]string{}

	for _, t := range tests {
		n := possible(t)
		if curr, ok := poss[t.op.code]; ok {
			n = intersect(n, curr)
		}
		poss[t.op.code] = n
	}
	opCodes := getOpCodes(poss)

	program := parseOps(readProgram())

	fmt.Println(execute(program, opCodes))
}

func execute(os []*op, codes []string) registers {
	r := registers{0, 0, 0, 0}
	for _, o := range os {
		r = ops[codes[o.code]](o.in1, o.in2, o.out, r)
	}
	return r
}

type registers []int

type opFunc func(int, int, int, registers) registers

type op struct {
	code int
	in1  int
	in2  int
	out  int
}

type test struct {
	before registers
	op     *op
	after  registers
}

func getOpCodes(poss map[int][]string) []string {
	trans := map[int]map[string]bool{}
	for i, p := range poss {
		trans[i] = map[string]bool{}
		for _, v := range p {
			trans[i][v] = true
		}
	}
	ret := make([]string, len(poss))
	for len(trans) > 0 {
		var fk int
		var fv string
		for k, v := range trans {
			if len(v) == 1 {
				fk = k
				fv = selectKey(v)
				break
			}
		}

		for _, v := range trans {
			delete(v, fv)
		}

		delete(trans, fk)
		ret[fk] = fv
	}
	return ret
}

func selectKey(m map[string]bool) string {
	for k := range m {
		return k
	}
	panic("empty")
}

func intersect(ss1, ss2 []string) []string {
	m := map[string]bool{}
	for _, s := range ss1 {
		m[s] = true
	}
	ret := []string{}
	for _, s := range ss2 {
		if _, ok := m[s]; ok {
			ret = append(ret, s)
		}
	}
	return ret
}

func possible(t *test) []string {
	ret := []string{}
	for n, o := range ops {
		if o(t.op.in1, t.op.in2, t.op.out, t.before).equals(t.after) {
			ret = append(ret, n)
		}
	}
	return ret
}

func (r registers) copy() registers {
	ret := registers{}
	for _, i := range r {
		ret = append(ret, i)
	}
	return ret
}

func (r registers) equals(o registers) bool {
	for i, v := range r {
		if o[i] != v {
			return false
		}
	}
	return true
}

func parseTests(lines []string) []*test {
	tsts := []*test{}
	var working *test
	for _, line := range lines {
		if strings.HasPrefix(line, "Before:") {
			var r1, r2, r3, r4 int
			_, err := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &r1, &r2, &r3, &r4)
			if err != nil {
				panic(err)
			}
			working = &test{}
			working.before = registers{r1, r2, r3, r4}
		} else if strings.HasPrefix(line, "After:") {
			var r1, r2, r3, r4 int
			_, err := fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &r1, &r2, &r3, &r4)
			if err != nil {
				panic(err)
			}
			working.after = registers{r1, r2, r3, r4}
			tsts = append(tsts, working)
			working = nil
		} else if working != nil {
			var oc, in1, in2, out int
			_, err := fmt.Sscanf(line, "%d %d %d %d", &oc, &in1, &in2, &out)
			if err != nil {
				panic(err)
			}
			working.op = &op{
				code: oc,
				in1:  in1,
				in2:  in2,
				out:  out,
			}
		}
	}
	if working != nil {
		tsts = append(tsts, working)
	}
	return tsts
}

func parseOps(lines []string) []*op {
	os := []*op{}
	for _, line := range lines {
		var oc, in1, in2, out int
		fmt.Println(line)
		_, err := fmt.Sscanf(line, "%d %d %d %d", &oc, &in1, &in2, &out)
		if err != nil {
			panic(err)
		}
		os = append(os, &op{
			code: oc,
			in1:  in1,
			in2:  in2,
			out:  out,
		})
	}
	return os
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}

func readProgram() []string {
	data, err := ioutil.ReadFile("program.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}

var ops = map[string]opFunc{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,

	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,

	"setr": setr,
	"seti": seti,

	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,

	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

func addr(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] + r[b]
	return r
}

func addi(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] + b
	return r
}

func mulr(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] * r[b]
	return r
}

func muli(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] * b
	return r
}

func banr(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] & r[b]
	return r
}

func bani(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] & b
	return r
}

func borr(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] | r[b]
	return r
}

func bori(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a] | b
	return r
}

func setr(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = r[a]
	return r
}

func seti(a, b, c int, r registers) registers {
	r = r.copy()
	r[c] = a
	return r
}

func gtir(a, b, c int, r registers) registers {
	r = r.copy()
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtri(a, b, c int, r registers) registers {
	r = r.copy()
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func gtrr(a, b, c int, r registers) registers {
	r = r.copy()
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqir(a, b, c int, r registers) registers {
	r = r.copy()
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqri(a, b, c int, r registers) registers {
	r = r.copy()
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}

func eqrr(a, b, c int, r registers) registers {
	r = r.copy()
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
	return r
}
