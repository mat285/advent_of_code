package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	sum := 0
	seen := map[int]bool{0: true}
	for {
		for _, line := range lines {
			i, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				continue
			}
			sum += i
			fmt.Println(sum)
			if _, ok := seen[sum]; ok {
				fmt.Println("found", sum)
				return
			}
			seen[sum] = true
		}
	}

	// fmt.Println(sum)
}

func main1() {
	data, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	sum := 0
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		sum += i
	}

	fmt.Println(sum)
}
