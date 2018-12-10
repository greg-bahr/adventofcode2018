package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Dx, Dy int
}

func (p *Point) simPoint() {
	p.X += p.Dx
	p.Y += p.Dy
}

func (p *Point) reverseSim() {
	p.X -= p.Dx
	p.Y -= p.Dy
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day10/input.txt")
	check(err)

	re := regexp.MustCompile("-?\\d+")
	lines := strings.Split(string(file), "\r\n")

	points := make([]Point, len(lines))

	for i, line := range lines {
		nums := re.FindAllString(line, -1)

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		dx, _ := strconv.Atoi(nums[2])
		dy, _ := strconv.Atoi(nums[3])

		points[i] = Point{x, y, dx, dy}
	}

	minDistance := float64(999999)
	bestIteration := -1

	for i := 0; i < 11000; i++ {
		for j := 0; j < len(points); j++ {
			points[j].simPoint()
		}

		distance := distance(calculateBounds(points))

		if distance < minDistance {
			minDistance = distance
			bestIteration = i + 1
		}
	}

	for i := 0; i < 11000-bestIteration; i++ {
		for j := 0; j < len(points); j++ {
			points[j].reverseSim()
		}
	}

	printArr(points)
	fmt.Println(bestIteration)
}

func printArr(points []Point) {
	upperLeft, lowerRight := calculateBounds(points)

	arr := make([][]string, lowerRight.Y-upperLeft.Y+1)
	for i := range arr {
		arr[i] = make([]string, lowerRight.X-upperLeft.X+1)
		for j := range arr[i] {
			arr[i][j] = " "
		}
	}

	for _, point := range points {
		arr[(upperLeft.Y*-1)+point.Y][(upperLeft.X*-1)+point.X] = "#"
	}

	for _, line := range arr {
		fmt.Println(line)
	}
}

func calculateBounds(points []Point) (upperLeft Point, lowerRight Point) {
	upperLeft = points[0]
	lowerRight = points[0]

	for _, point := range points {
		if point.X < upperLeft.X {
			upperLeft.X = point.X
		} else if point.X > lowerRight.X {
			lowerRight.X = point.X
		}
		if point.Y < upperLeft.Y {
			upperLeft.Y = point.Y
		} else if point.Y > lowerRight.Y {
			lowerRight.Y = point.Y
		}
	}

	return upperLeft, lowerRight
}

func distance(point1, point2 Point) float64 {
	return math.Sqrt(math.Pow(float64(point1.X-point2.X), 2) + math.Pow(float64(point1.Y-point2.Y), 2))
}
