package main

import (
	"fmt"
	"math"
)

func main() {
	serial := 8199
	grid := [300][300]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = calculatePowerLevel(j+1, i+1, serial)
		}
	}

	maxPower := -99999
	x := 0
	y := 0

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			power := calculateSquarePower(j, i, 3, grid)

			if power > maxPower {
				maxPower = power
				x = j + 1
				y = i + 1
			}
		}
	}

	fmt.Println(x, y)

	size := -1
	maxPower = -999999
	x = -1
	y = -1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			bestPower, bestSize := largestTotalSquare(j, i, grid)

			if bestPower > maxPower {
				maxPower = bestPower
				size = bestSize
				x = j + 1
				y = i + 1
			}
		}
	}

	fmt.Println(x, y, size)
}

func calculatePowerLevel(x, y, serial int) int {
	rackId := x + 10
	power := ((rackId * y) + serial) * rackId

	return ((power / 100) % 10) - 5
}

func calculateSquarePower(x, y, size int, grid [300][300]int) int {
	sum := 0

	for i := y; i < y+size; i++ {
		for j := x; j < x+size; j++ {
			sum += grid[i][j]
		}
	}

	return sum
}

func maxSquareSize(x, y int) int {
	return 300 - int(math.Max(float64(x), float64(y)))
}

func largestTotalSquare(x, y int, grid [300][300]int) (int, int) {
	max := -99999
	size := -1
	maxSize := maxSquareSize(x, y)

	if maxSize < 3 || maxSize > 50 {
		return max, size
	}

	for i := 1; i < maxSize; i++ {
		power := calculateSquarePower(x, y, i, grid)

		//fmt.Println(x, y, power, i)

		if power > max {
			max = power
			size = i
		}
	}

	return max, size
}
