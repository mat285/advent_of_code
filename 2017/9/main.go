package main

import (
	"fmt"
	"io/ioutil"
)

var (
	test1 = []byte("{{{},{},{{}}}}")
)

func main() {
	bs := read() //test1
	part1(bs)
	part2(bs)
}

func part1(bs []byte) {
	group, _ := process(bs)
	fmt.Println(group.score())
}

func part2(bs []byte) {
	_, num := process(bs)
	fmt.Println(num)
}

type stack struct {
	head *node
	size int
}

type node struct {
	val  *byteOrGroup
	next *node
}

type group struct {
	str string
	sub []*group
}

type byteOrGroup struct {
	b byte
	g *group
}

type stringOrGroup struct {
	s string
	g *group
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(v *byteOrGroup) {
	n := &node{val: v}
	if s.head != nil {
		n.next = s.head
	}
	s.size++
	s.head = n
}

func (s *stack) pop() *byteOrGroup {
	if s.head == nil {
		panic("pop empty stack")
	}
	n := s.head
	s.head = s.head.next
	s.size--
	return n.val
}

func fromByte(b byte) *byteOrGroup {
	return &byteOrGroup{
		b: b,
	}
}

func fromGroup(g *group) *byteOrGroup {
	return &byteOrGroup{
		g: g,
	}
}

func (g *group) score(depth ...int) int {
	sum := 1
	start := 0
	if len(depth) > 0 {
		sum += depth[0]
		start = depth[0]
	}
	for _, c := range g.sub {
		sum += c.score(start + 1)
	}
	return sum
}

func process(bs []byte) (*group, int) {
	stack := newStack()
	state := 0
	count := 0
	for i := 0; i < len(bs); i++ {
		b := bs[i]
		switch state {
		case 0:
			if b == '!' {
				state = 2
			} else if b == ',' {
				continue
			} else if b == '<' {
				state = 1
			} else if b != '}' {
				stack.push(fromByte(b))
			} else {
				str := ""
				gs := []*group{}
				for stack.head.val.b != '{' {
					bg := stack.pop()
					if bg.g != nil {
						gs = append(gs, bg.g)
					} else {
						str += string(bg.b)
					}
				}
				stack.pop()
				stack.push(fromGroup(&group{
					str: str,
					sub: gs,
				}))
			}
		case 1:
			if b == '>' {
				state = 0
			} else if b == '!' {
				state = 3
			} else {
				count++
			}
		case 2:
		case 3:
			state -= 2
		}
	}
	if stack.head.next != nil {
		fmt.Println(stack.size)
		fmt.Println(string(stack.head.next.val.b))
		panic("bad stack")
	}

	return stack.head.val.g, count
}

func read() []byte {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return data[:len(data)-1] // trim newline
}
