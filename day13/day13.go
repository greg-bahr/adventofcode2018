package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Train struct {
	Position  [2]int
	Direction int
	LastTurn  int
	Crashed   bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	DOWN  = iota
	LEFT  = iota
	UP    = iota
	RIGHT = iota
)

func main() {
	file, err := ioutil.ReadFile("./day13/input.txt")
	check(err)

	lines := strings.Split(string(file), "\r\n")

	var trains []Train

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if string(lines[i][j]) == "v" {
				trains = append(trains, Train{[2]int{i, j}, DOWN, RIGHT, false})
			} else if string(lines[i][j]) == ">" {
				trains = append(trains, Train{[2]int{i, j}, RIGHT, RIGHT, false})
			} else if string(lines[i][j]) == "<" {
				trains = append(trains, Train{[2]int{i, j}, LEFT, RIGHT, false})
			} else if string(lines[i][j]) == "^" {
				trains = append(trains, Train{[2]int{i, j}, UP, RIGHT, false})
			}
		}
	}

	firstCrash := [2]int{-1, -1}
	var notCrashed Train
	numCrashed := 0

	for len(trains)-numCrashed != 1 {
		numCrashed = 0

		sort.Slice(trains, func(i, j int) bool {
			return (trains[i].Position[0] < trains[j].Position[0]) || (trains[i].Position[1] < trains[j].Position[1])
		})

		for i := 0; i < len(trains); i++ {
			if !trains[i].Crashed {
				trains[i] = newLocation(trains[i], lines)
			}

			if !trains[i].Crashed && crashed(trains[i], trains) {
				trains[i].Crashed = true

				if firstCrash[0] == -1 {
					firstCrash = trains[i].Position
				}
			}
		}

		for _, train := range trains {
			if train.Crashed {
				numCrashed++
			} else {
				notCrashed = train
			}
		}
	}

	fmt.Println(firstCrash[1], firstCrash[0])
	fmt.Println(notCrashed.Position[1], notCrashed.Position[0])
}

func newLocation(train Train, input []string) Train {
	x := train.Position[1]
	y := train.Position[0]

	switch train.Direction {
	case UP:
		train.Position[0] = y - 1
		y -= 1
	case DOWN:
		train.Position[0] = y + 1
		y += 1
	case LEFT:
		train.Position[1] = x - 1
		x -= 1
	case RIGHT:
		train.Position[1] = x + 1
		x += 1
	}

	if string(input[y][x]) == "/" {
		switch train.Direction {
		case UP:
			train.Direction = RIGHT
		case DOWN:
			train.Direction = LEFT
		case LEFT:
			train.Direction = DOWN
		case RIGHT:
			train.Direction = UP
		}
	} else if string(input[y][x]) == "\\" {
		switch train.Direction {
		case UP:
			train.Direction = LEFT
		case DOWN:
			train.Direction = RIGHT
		case LEFT:
			train.Direction = UP
		case RIGHT:
			train.Direction = DOWN
		}
	} else if string(input[y][x]) == "+" {
		switch train.LastTurn {
		case RIGHT:
			train.Direction = (train.Direction - 1) % 4
			if train.Direction < 0 {
				train.Direction = 3
			}
			train.LastTurn = LEFT
		case LEFT:
			train.LastTurn = UP
		case UP:
			train.Direction = (train.Direction + 1) % 4
			train.LastTurn = RIGHT
		}
	}

	return train
}

func crashed(train Train, trains []Train) bool {
	for i := 0; i < len(trains); i++ {
		if train != trains[i] && train.Position == trains[i].Position && !trains[i].Crashed {
			trains[i].Crashed = true

			return true
		}
	}

	return false
}
