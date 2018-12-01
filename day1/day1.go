package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day1/input.txt")
	check(err)

	nums := strings.Split(string(file), "\n")

	fmt.Printf("Part 1: %d\n", part1(nums))
	fmt.Printf("Part 2: %d\n", part2(nums))
}

func part1(nums []string) int {
	frequency := 0

	for i := 0; i < len(nums); i++ {
		num, _ := strconv.Atoi(strings.TrimSpace(nums[i]))
		frequency += num
	}

	return frequency
}

func part2(nums []string) int {
	frequency := 0
	set := make(map[int]bool)

	for {
		for i := 0; i < len(nums); i++ {
			num, _ := strconv.Atoi(strings.TrimSpace(nums[i]))
			frequency += num

			if set[frequency] {
				return frequency
			}

			set[frequency] = true
		}
	}
}