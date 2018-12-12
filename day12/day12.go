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
	file, err := ioutil.ReadFile("./day12/input.txt")
	check(err)

	lines := strings.Split(string(file), "\r\n")

	currentState := make(map[int]int)
	possibleStates := make(map[int]int)

	initialState := strings.TrimSpace(strings.Split(lines[0], " ")[2])

	for i, char := range initialState {
		if string(char) == "#" {
			currentState[i] = 1
		} else {
			currentState[i] = 0
		}
	}

	lines = lines[2:]
	for _, line := range lines {
		line = strings.Replace(line, "#", "1", -1)
		line = strings.Replace(line, ".", "0", -1)

		arr := strings.Split(line, " => ")
		before, _ := strconv.Atoi(arr[0])
		after, _ := strconv.Atoi(arr[1])

		possibleStates[before] = after
	}

	lastSum := sumPots(currentState)

	for i := 0; ; i++ {
		newState := copyState(currentState)

		for i := minPot(currentState) - 3; i < maxPot(currentState)+3; i++ {
			newState[i] = getPotStatus(i, currentState, possibleStates)
		}
		currentState = newState
		currentSum := sumPots(currentState)

		if i == 19 {
			fmt.Println(sumPots(currentState))
		}

		if currentSum-lastSum == 23 {
			fmt.Println(((50000000000 - i) * 23) + lastSum)
			break
		}

		lastSum = currentSum
	}
}

func getPotStatus(pot int, currentState, possibleStates map[int]int) int {
	num := 0

	num += currentState[pot-2] * 10000
	num += currentState[pot-1] * 1000
	num += currentState[pot] * 100
	num += currentState[pot+1] * 10
	num += currentState[pot+2]

	return possibleStates[num]
}

func copyState(state map[int]int) map[int]int {
	newState := make(map[int]int)

	for key, value := range state {
		newState[key] = value
	}

	return newState
}

func minPot(pots map[int]int) int {
	min := 9999

	for key := range pots {
		if key < min {
			min = key
		}
	}

	return min
}

func maxPot(pots map[int]int) int {
	max := -9999

	for key := range pots {
		if key > max {
			max = key
		}
	}

	return max
}

func printCurrentState(pots map[int]int) {
	for i := minPot(pots) - 3; i < maxPot(pots)+3; i++ {
		if pots[i] == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func sumPots(pots map[int]int) int {
	sum := 0
	for pot, val := range pots {
		if val == 1 {
			sum += pot
		}
	}

	return sum
}
