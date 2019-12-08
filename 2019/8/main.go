package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var (
	width  = 25
	height = 6
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	part1(input)
	part2(input)
}

func part1(input string) {
	image := parse(input, width, height)

	minL := [][]int{}
	min := math.MaxInt32
	for _, layer := range image {
		c := count(layer, 0)
		if c < min {
			min = c
			minL = layer
		}
	}
	fmt.Println(count(minL, 1) * count(minL, 2))
}

func part2(input string) {
	image := parse(input, width, height)
	final := render(image)

	for _, row := range final {
		for _, v := range row {
			if v == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func render(image [][][]int) [][]int {
	final := make([][]int, len(image[0]))
	for i := 0; i < len(final); i++ {
		final[i] = make([]int, len(image[0][0]))
		for j := 0; j < len(final[i]); j++ {
			final[i][j] = 2
		}
	}

	for z := 0; z < len(image); z++ {
		for y := 0; y < len(image[z]); y++ {
			for x := 0; x < len(image[z][y]); x++ {
				if final[y][x] == 2 {
					final[y][x] = image[z][y][x]
				}
			}
		}
	}
	return final
}

func count(layer [][]int, i int) int {
	count := 0
	for _, row := range layer {
		for _, v := range row {
			if v == i {
				count++
			}
		}
	}
	return count
}

func parse(input string, w, h int) [][][]int {
	layers := make([][][]int, len(input)/(w*h))
	for z := 0; z < len(layers); z++ {
		layers[z] = make([][]int, h)
		for y := 0; y < len(layers[z]); y++ {
			layers[z][y] = make([]int, w)
		}
	}
	var x, y, z int
	for _, c := range input {
		d, _ := strconv.Atoi(string(c))
		layers[z][y][x] = d
		x++
		if x == w {
			x = 0
			y++
		}
		if y == h {
			y = 0
			z++
		}
	}
	return layers
}
