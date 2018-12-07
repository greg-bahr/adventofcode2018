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

	points := make(map[Point]int)

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

	for point, _ := range points {
		area := 1

		queue := []Point{}
		visited := make(map[Point]bool)

		queue = append(queue, point)

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			for _, neighbor := range findNeighbors(current) {
				parent, multipleParents := closestParent(neighbor, points)

				if !visited[neighbor] && !queueContains(neighbor, queue) && !multipleParents && parent == point {
					if inBounds(neighbor, upperLeftBound, lowerRightBound) {
						area++
						queue = append(queue, neighbor)
					} else {
						queue = []Point{}
						area = -1
						break
					}
				}
			}

			visited[current] = true
		}

		points[point] = area
	}

	maxArea := 0
	for _, area := range points {
		if area > maxArea {
			maxArea = area
		}
	}
	fmt.Println(maxArea)

	var safePoint Point
	for point := range points {
		if isSafe(point, points) {
			safePoint = point
			break
		}
	}

	area := 1
	queue := []Point{}
	visited := make(map[Point]bool)

	queue = append(queue, safePoint)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range findNeighbors(current) {
			if !visited[neighbor] && !queueContains(neighbor, queue) && isSafe(neighbor, points) {
				area++
				queue = append(queue, neighbor)
			}
		}

		visited[current] = true
	}

	fmt.Println(area)
}

func inBounds(point Point, upperLeft Point, lowerRight Point) bool {
	return point.X >= upperLeft.X && point.X <= lowerRight.X && point.Y >= upperLeft.Y && point.Y <= lowerRight.Y
}

func distance(point1 Point, point2 Point) int {
	return int(math.Abs(float64(point1.X-point2.X)) + math.Abs(float64(point1.Y-point2.Y)))
}

func findNeighbors(point Point) []Point {
	return []Point{
		{point.X + 1, point.Y},
		{point.X, point.Y + 1},
		{point.X - 1, point.Y},
		{point.X, point.Y - 1},
		{point.X - 1, point.Y - 1},
		{point.X + 1, point.Y - 1},
		{point.X + 1, point.Y + 1},
		{point.X - 1, point.Y + 1},
	}
}

func queueContains(point Point, queue []Point) bool {
	for _, item := range queue {
		if point == item {
			return true
		}
	}

	return false
}

func closestParent(point Point, points map[Point]int) (Point, bool) {
	var minParent Point
	minDistance := 99999
	nextMinDistance := 99999

	for parent := range points {
		if distance(point, parent) <= minDistance {
			minParent = parent
			nextMinDistance = minDistance
			minDistance = distance(point, parent)
		}
	}

	return minParent, minDistance == nextMinDistance
}

func isSafe(point Point, points map[Point]int) bool {
	total := 0

	for parent := range points {
		total += distance(point, parent)
	}

	return total < 10000
}
