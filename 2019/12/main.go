package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	moons := parse(lines)
	simulate(moons, 1000, false)
	fmt.Println("Energy:", energy(moons))
}

func part2(lines []string) {
	cycles := point{}
	for i := 0; i < len(cycles); i++ {
		cycles[i] = cycle(parse(lines), i, false)
	}
	fmt.Println("Cycle length:", lcm(cycles[:]...))
}

type point [3]int

const (
	x = 0
	y = 1
	z = 2
)

type moon struct {
	name string
	pos  point
	vel  point
}

func cycle(moons []*moon, i int, debug bool) int {
	initial := state(moons, i)
	steps := 0
	for {
		steps++
		simulate(moons, 1, debug)
		state := state(moons, i)
		if state == initial {
			return steps
		}
	}
}

func simulate(moons []*moon, ticks int, debug bool) {
	for i := 0; i < ticks; i++ {
		applyGravity(moons)
		applyVelocity(moons)
		if debug {
			fmt.Printf("After %d steps:\n", i+1)
			for _, m := range moons {
				fmt.Println(m)
			}
			fmt.Println()
		}
	}
}

func state(moons []*moon, i int) string {
	state := ""
	forEach(moons, func(m *moon) { state += fmt.Sprintf("p:%d,v:%d", m.pos[i], m.vel[i]) })
	return state
}

func energy(moons []*moon) int {
	sum := 0
	forEach(moons, func(m *moon) { sum += m.energy() })
	return sum
}

func applyGravity(moons []*moon) {
	forEach(moons, func(m *moon) {
		forEach(moons, func(o *moon) {
			m.applyGravity(o)
		})
	})
}

func applyVelocity(moons []*moon) {
	forEach(moons, func(m *moon) { m.applyVelocity() })
}

func (m *moon) applyVelocity() {
	m.pos = add(m.pos, m.vel)
}

func (m *moon) applyGravity(o *moon) {
	for i := 0; i < len(m.vel); i++ {
		m.vel[i] += pullTo(m.pos[i], o.pos[i])
	}
}

func pullTo(src, dst int) int {
	if src == dst {
		return 0
	}
	return (dst - src) / abs(dst-src)
}

func (m *moon) potential() int {
	return m.pos.absSum()
}

func (m *moon) kinetic() int {
	return m.vel.absSum()
}

func (m *moon) energy() int {
	return m.potential() * m.kinetic()
}

func (m *moon) String() string {
	return fmt.Sprintf("pos=%s, vel=%s, %s", m.pos, m.vel, m.name)
}

func (p point) String() string {
	return fmt.Sprintf("<x= %d, y= %d, z= %d>", p[x], p[y], p[z])
}

func (p point) absSum() int {
	sum := 0
	for i := 0; i < len(p); i++ {
		sum += abs(p[i])
	}
	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(is ...int) int {
	mul := 1
	for _, i := range is {
		mul = (mul / gcd(mul, i)) * i
	}
	return mul
}

func add(p, q point) point {
	ret := point{}
	for i := 0; i < len(p); i++ {
		ret[i] = p[i] + q[i]
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func forEach(moons []*moon, fn func(*moon)) {
	for _, m := range moons {
		fn(m)
	}
}

func parse(lines []string) []*moon {
	names := []string{"Io", "Europa", "Ganymede", "Callisto"}
	moons := make([]*moon, len(names))
	if len(lines) != len(names) {
		panic("Wrong number of moons")
	}

	for i := 0; i < len(names); i++ {
		pos := point{}
		_, err := fmt.Sscanf(lines[i], "<x=%d, y=%d, z=%d>", &pos[x], &pos[y], &pos[z])
		if err != nil {
			panic(err)
		}
		moons[i] = &moon{
			name: names[i],
			pos:  pos,
			vel:  point{},
		}
	}
	return moons
}
