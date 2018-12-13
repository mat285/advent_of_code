package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var testLines = strings.Split(
	`/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `, "\n")

var fun = strings.Split(
	`/->-\
| /-+-\
| | | |
\-+-/ |
  \---/`, "\n",
)

const (
	up    = "^"
	down  = "v"
	left  = "<"
	right = ">"

	vert  = "|"
	horz  = "-"
	crdr  = "\\"
	crur  = "/"
	inter = "+"
)

func main() {
	lines := read()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	ts, cs := parse(lines)
	var coll *point
	printT(ts, cs, nil)
	count := 0
	for coll == nil {
		sort.Slice(cs, func(i, j int) bool {
			return less(cs[i], cs[j])
		})
		for _, c := range cs {
			c.move(ts)
			coll = collision(c, cs)
			if coll != nil {
				break
			}
		}
		count++
	}
	fmt.Println(coll)
}

func part2(lines []string) {
	ts, cs := parse(lines)

	for len(cs) > 1 {
		sort.Slice(cs, func(i, j int) bool {
			return less(cs[i], cs[j])
		})

		for i := 0; i < len(cs); i++ {
			c := cs[i]
			c.move(ts)
			coll := collision(c, cs)
			if coll != nil {
				n := []*cart{}
				for _, a := range cs {
					if !(coll.x == a.pos.x && coll.y == a.pos.y) {
						n = append(n, a)
					}
				}
				cs = n
			}
		}
	}
	fmt.Println(cs[0].pos)
}

type point struct{ x, y int }

type track struct {
	v   string
	pos point
}

type cart struct {
	id        int
	pos       point
	direction string
	prev      string
}

func less(c1, c2 *cart) bool {
	if c1.pos.y < c2.pos.y {
		return true
	} else if c1.pos.y == c2.pos.y {
		return c1.pos.x < c2.pos.x
	}
	return false
}

func printT(ts [][]*track, cs []*cart, coll *point) {
	for i, t := range ts {
		for j, a := range t {
			printed := false
			if coll != nil && coll.x == j && coll.y == i {
				fmt.Print("x")
				continue
			}

			for _, c := range cs {
				if c.pos.x == j && c.pos.y == i {
					fmt.Print(c.direction)
					printed = true
				}
			}
			if a != nil && !printed {
				fmt.Print(a.v)
			} else if printed {
				continue
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (c *cart) move(ts [][]*track) {
	// update pos
	switch c.direction {
	case up:
		c.pos.y--
	case down:
		c.pos.y++
	case left:
		c.pos.x--
	case right:
		c.pos.x++
	default:
		panic("unknown d")
	}
	// set direction
	switch ts[c.pos.y][c.pos.x].v {
	case vert:
		break
	case horz:
		break
	case crdr:
		if c.direction == up {
			c.direction = left
		} else if c.direction == down {
			c.direction = right
		} else if c.direction == left {
			c.direction = up
		} else if c.direction == right {
			c.direction = down
		} else {
			panic("uh oh spaghettios")
		}
	case crur:
		if c.direction == down {
			c.direction = left
		} else if c.direction == up {
			c.direction = right
		} else if c.direction == right {
			c.direction = up
		} else if c.direction == left {
			c.direction = down
		} else {
			panic("some t not goos")
		}
	case inter:
		d := c.turn()
		switch c.direction {
		case left:
			if d == left {
				c.direction = down
			} else if d == right {
				c.direction = up
			}
		case right:
			if d == left {
				c.direction = up
			} else if d == right {
				c.direction = down
			}
		case up:
			if d == left || d == right {
				c.direction = d
			}
		case down:
			if d == left || d == right {
				c.direction = reverseD(d)
			}
		default:
			panic("nooooo")
		}
	default:
		panic("unkown tile")
	}
}

func (c *cart) turn() string {
	if c.prev == "" || c.prev == right {
		c.prev = left
	} else if c.prev == left {
		c.prev = vert
	} else {
		c.prev = right
	}
	return c.prev
}

func collision(c *cart, cs []*cart) *point {
	for i := 0; i < len(cs)-1; i++ {
		if c.id != cs[i].id && c.pos.x == cs[i].pos.x && c.pos.y == cs[i].pos.y {
			return &point{c.pos.x, c.pos.y}
		}
	}
	return nil
}

func parse(lines []string) ([][]*track, []*cart) {
	ts := make([][]*track, len(lines))
	cs := []*cart{}
	for i, line := range lines {
		bs := []byte(line)
		ts[i] = make([]*track, len(bs))
		for j, b := range bs {
			if b == ' ' {
				continue
			} else if isCart(b) {
				cs = append(cs, &cart{
					id:        len(cs),
					pos:       point{j, i},
					direction: string(b),
				})
				if string(b) == up || string(b) == down {
					ts[i][j] = &track{
						v:   vert,
						pos: point{j, i},
					}
				} else {
					ts[i][j] = &track{
						v:   horz,
						pos: point{j, i},
					}
				}
			} else {
				ts[i][j] = &track{
					v:   string(b),
					pos: point{j, i},
				}
			}
		}
	}
	return ts, cs
}

func reverseD(dir string) string {
	switch dir {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	default:
		panic("oh well hello there")
	}
}

func isCart(b byte) bool {
	s := string(b)
	return s == up || s == down || s == left || s == right
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}
