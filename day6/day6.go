package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day6/input.txt")
	check(err)

	re := regexp.MustCompile("\\d+")
	lines := strings.Split(string(file), "\r\n")

	points := map[Point]int{}

	for _, line := range lines {
		nums := re.FindAllString(line, -1)

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		points[Point{x, y}] = 0
	}

	upperLeftBound := Point{321, 199}
	lowerRightBound := Point{321, 199}

	for point, _ := range points {
		if point.X < upperLeftBound.X {
			upperLeftBound.X = point.X
		} else if point.X > lowerRightBound.X {
			lowerRightBound.X = point.X
		}

		if point.Y < upperLeftBound.Y {
			upperLeftBound.Y = point.Y
		} else if point.Y > lowerRightBound.Y {
			lowerRightBound.Y = point.Y
		}
	}

	fmt.Println(upperLeftBound)
	fmt.Println(lowerRightBound)

	fmt.Println(inBounds(Point{321, 199}, upperLeftBound, lowerRightBound))
}

func inBounds(point Point, upperLeft Point, lowerRight Point) bool {
	return point.X >= upperLeft.X && point.X <= lowerRight.X && point.Y >= upperLeft.Y && point.Y <= lowerRight.Y
}
