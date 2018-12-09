package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day9/input.txt")
	check(err)

	re := regexp.MustCompile("\\d+")
	nums := re.FindAllString(string(file), -1)

	numPlayers, _ := strconv.Atoi(nums[0])
	lastWorth, _ := strconv.Atoi(nums[1])

	fmt.Println(calculateWinningScore(numPlayers, lastWorth))
	fmt.Println(calculateWinningScore(numPlayers, lastWorth*100))
}

func calculateWinningScore(numPlayers int, lastWorth int) int {
	marbleValue := 1
	currentPlayer := 0
	marbles := ring.New(1)
	players := make([]int, numPlayers)
	marbles.Value = 0

	for marbleValue <= lastWorth {
		if marbleValue%23 == 0 {
			players[currentPlayer] += marbleValue

			for i := 0; i < 8; i++ {
				marbles = marbles.Prev()
			}
			players[currentPlayer] += marbles.Next().Value.(int)
			marbles.Unlink(1)
			marbles = marbles.Next()
		} else {
			marbles = marbles.Next()
			marbles.Link(&ring.Ring{Value: marbleValue})
			marbles = marbles.Next()
		}

		marbleValue++
		currentPlayer = (currentPlayer + 1) % numPlayers
	}

	maxScore := 0
	for _, score := range players {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}
