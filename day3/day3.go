package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day3/input.txt")
	check(err)

	lines := strings.Split(string(file), "\r\n")

	fabric := [1000][1000]int{}
	sum := 0

	re := regexp.MustCompile("[0-9]+")

	for _, line := range lines {
		nums := re.FindAllString(line, -1)

		id, _ := strconv.Atoi(nums[0])
		left, _ := strconv.Atoi(nums[1])
		top, _ := strconv.Atoi(nums[2])
		width, _ := strconv.Atoi(nums[3])
		height, _ := strconv.Atoi(nums[4])

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if fabric[i+top][j+left] > 0 {
					fabric[i+top][j+left] = -1
				} else if fabric[i+top][j+left] == 0 {
					fabric[i+top][j+left] = id
				}
			}
		}
	}

	for _, line := range fabric {
		for _, num := range line {
			if num == -1 {
				sum++
			}
		}
	}

	fmt.Println(sum)

	for _, line := range lines {
		nums := re.FindAllString(line, -1)

		id, _ := strconv.Atoi(nums[0])
		left, _ := strconv.Atoi(nums[1])
		top, _ := strconv.Atoi(nums[2])
		width, _ := strconv.Atoi(nums[3])
		height, _ := strconv.Atoi(nums[4])

		overlaps := false

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if fabric[i+top][j+left] == -1 {
					overlaps = true
				}
			}
		}

		if !overlaps {
			fmt.Println(id)
			return
		}
	}
}
