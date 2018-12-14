package main

import (
	"encoding/hex"
	"fmt"
)

var lengths = []byte{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118}
var extra = []byte{17, 31, 73, 47, 23}

func main() {
	part1(duplicate(lengths))
	part2()
}

func part1(input []byte) {
	r := newRing(iter(256))
	r.round(input)
	fmt.Println(int(r.buf[0]) * int(r.buf[1]))
}

func part2() {
	str := asString(lengths)
	fmt.Println(str)
	fmt.Println(Hash(str))
}

func asString(input []byte) string {
	s := ""
	for _, i := range input {
		s += fmt.Sprintf("%d,", i)
	}
	return s[:len(s)-1]
}

type ring struct {
	buf      []byte
	pos      int
	skipSize int
}

func Hash(s string) string {
	r := newRing(iter(256))
	ls := []byte(s)
	ls = append(ls, extra...)
	for i := 0; i < 64; i++ {
		r.round(ls)
	}
	return hex.EncodeToString(r.dense())
}

func newRing(buf []byte) *ring {
	return &ring{
		buf:      buf,
		pos:      0,
		skipSize: 0,
	}
}

func (r *ring) dense() []byte {
	ret := make([]byte, 16)
	for i := 0; i < len(r.buf); i++ {
		idx := i / 16
		if i%16 == 0 {
			ret[idx] = r.buf[i]
		} else {
			ret[idx] ^= r.buf[i]
		}
	}
	return ret
}

func (r *ring) round(ls []byte) {
	for _, l := range ls {
		r.act(l)
	}
}

func (r *ring) skip(i int) {
	r.pos = (r.pos + i) % len(r.buf)
}

func (r *ring) reverse(l byte) {
	for i := 0; i < int(l)/2; i++ {
		start := (r.pos + i) % len(r.buf)
		end := (r.pos + int(l) - i - 1) % len(r.buf)
		v1 := r.buf[start]
		r.buf[start] = r.buf[end]
		r.buf[end] = v1
	}
}

func (r *ring) act(l byte) {
	r.reverse(l)
	r.skip(int(l) + r.skipSize)
	r.skipSize++
}

func iter(l int) []byte {
	ret := make([]byte, l)
	for i := 0; i < l; i++ {
		ret[i] = byte(i)
	}
	return ret
}

func duplicate(i []byte) []byte {
	dst := make([]byte, len(i))
	copy(dst, i)
	return dst
}
