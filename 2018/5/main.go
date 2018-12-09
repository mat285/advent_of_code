package main

import (
	"fmt"
	"io/ioutil"
)

const (
	polar = 32
)

func main() {
	part1()
	part2()
}

func part1() {
	b := read()
	l := toList(b)
	size := len(b)
	l, size = pass(l, size)
	// fmt.Println(string(toSlice(l, size)))
	fmt.Println(size)
}

func part2() {
	b := read()
	min := len(b)
	for i := byte('a'); i <= 'z'; i++ {
		l := toList(b)
		size := len(b)
		l, size = remove(l, size, i)
		l, size = pass(l, size)
		if size < min {
			min = size
		}
	}
	fmt.Println(min)
}

type node struct {
	val  byte
	next *node
	prev *node
}

func pass(list *node, size int) (*node, int) {
	if list == nil || list.next == nil {
		return list, size
	}

	head := &node{}
	head.next = list
	list.prev = head

	curr := head.next
	for curr.next != nil {
		if diff(curr.val, curr.next.val) == polar {
			t := curr.next.next
			if t != nil {
				t.prev = curr.prev
				t.prev.next = t
				curr = t.prev
			} else {
				curr.prev.next = nil
				curr = curr.prev
			}
			size -= 2
			if curr == head {
				curr = curr.next
			}
		} else {
			curr = curr.next
		}
	}
	head.next.prev = nil
	return head.next, size
}

func remove(list *node, size int, rm byte) (*node, int) {
	head := &node{}
	head.next = list
	list.prev = head

	curr := head.next
	for curr != nil {
		if diff(curr.val, rm) == 0 || diff(curr.val, rm) == polar {
			curr.prev.next = curr.next
			if curr.next != nil {
				curr.next.prev = curr.prev
			}
			size--
		}
		curr = curr.next
	}

	head.next.prev = nil
	return head.next, size
}

func toSlice(head *node, size int) []byte {
	s := make([]byte, 0, size)
	for head != nil {
		s = append(s, head.val)
		head = head.next
	}
	return s
}

func toList(bytes []byte) *node {
	head := &node{}
	head.val = bytes[0]

	curr := head

	for i := 1; i < len(bytes); i++ {
		n := &node{}
		n.val = bytes[i]
		n.prev = curr
		curr.next = n
		curr = n
	}
	return head
}

func diff(a, b byte) byte {
	if a >= b {
		return a - b
	}
	return b - a
}

func read() []byte {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return data[:len(data)-1]
}
