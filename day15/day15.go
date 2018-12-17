package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Point struct {
	X, Y int
}

type Unit struct {
	Id          int
	Name        string
	Position    Point
	AttackPower int
	Health      int
}

func main() {
	file, _ := ioutil.ReadFile("./day15/test.txt")

	input := strings.Split(string(file), "\r\n")
	var lines [][]string

	for _, line := range input {
		lines = append(lines, strings.Split(line, ""))
	}

	id := 0
	rounds := 0
	var goblins []*Unit
	var elves []*Unit
	var units []*Unit

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if string(lines[i][j]) == "E" {
				newUnit := Unit{id, "Elf", Point{j, i}, 3, 200}
				units = append(units, &newUnit)
				elves = append(elves, &newUnit)
				id++
			} else if string(lines[i][j]) == "G" {
				newUnit := Unit{id, "Goblin", Point{j, i}, 3, 200}
				units = append(units, &newUnit)
				goblins = append(goblins, &newUnit)
				id++
			}
		}
	}

	for hasAlive(elves) && hasAlive(goblins) {
		sortUnits(units)

		for _, unit := range units {
			if unit.Health < 1 {
				continue
			}
			if !hasAlive(elves) || !hasAlive(goblins) {
				sum := winningSum(units)

				fmt.Println(sum, rounds)
				fmt.Println(sum * rounds)
				return
			}

			adjacentTargets := canAttack(unit, units)
			if len(adjacentTargets) == 0 {
				var targetPoints []Point
				if unit.Name == "Elf" {
					targetPoints = getTargetPoints(goblins, lines)
				} else {
					targetPoints = getTargetPoints(elves, lines)
				}

				var moves []Point
				minPath := 99999
				for _, target := range targetPoints {
					path := pathSearch(unit.Position, target, lines)

					if len(path) == minPath {
						moves = append(moves, path[0])
					} else if len(path) > 0 && len(path) < minPath {
						moves = []Point{path[0]}
						minPath = len(path)
					}
				}

				sort.Slice(moves, func(i, j int) bool {
					return moves[i].X <= moves[j].X && moves[i].Y <= moves[j].Y
				})

				if len(moves) > 0 {
					lines[unit.Position.Y][unit.Position.X] = "."
					unit.Position = moves[0]

					if unit.Name == "Goblin" {
						lines[unit.Position.Y][unit.Position.X] = "G"
					} else {
						lines[unit.Position.Y][unit.Position.X] = "E"
					}
				}
			}

			adjacentTargets = canAttack(unit, units)
			if len(adjacentTargets) > 0 {
				var minUnit *Unit
				for _, point := range adjacentTargets {
					unit := getUnit(point, units)

					if minUnit == nil {
						minUnit = unit
					} else if unit.Health < minUnit.Health {
						minUnit = unit
					}
				}

				minUnit.Health -= unit.AttackPower
				if minUnit.Health < 1 {
					lines[minUnit.Position.Y][minUnit.Position.X] = "."
					minUnit.Position = Point{-1, -1}
				}
			}
		}

		fmt.Println("Round", rounds+1)
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println()

		rounds++
	}

	sum := winningSum(units)

	fmt.Println(sum, rounds)
	fmt.Println(sum * rounds)
}

func canAttack(unit *Unit, units []*Unit) []Point {
	var enemies []Point

	for _, diff := range []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		neighbor := Point{unit.Position.X + diff.X, unit.Position.Y + diff.Y}
		neighborUnit := getUnit(neighbor, units)

		if neighborUnit != nil && unit.Name == "Elf" && neighborUnit.Name == "Goblin" {
			enemies = append(enemies, neighbor)
		} else if neighborUnit != nil && unit.Name == "Goblin" && neighborUnit.Name == "Elf" {
			enemies = append(enemies, neighbor)
		}
	}

	return enemies
}

func hasAlive(units []*Unit) bool {
	for _, unit := range units {
		if unit.Health > 0 {
			return true
		}
	}

	return false
}

func getUnit(point Point, units []*Unit) *Unit {
	for _, unit := range units {
		if unit.Position == point {
			return unit
		}
	}

	return nil
}

func winningSum(units []*Unit) int {
	sum := 0

	for _, unit := range units {
		if unit.Health > 0 {
			sum += unit.Health
		}
	}

	return sum
}

func sortUnits(units []*Unit) {
	sort.Slice(units, func(i, j int) bool {
		return units[i].Position.Y <= units[j].Position.Y && units[i].Position.X <= units[j].Position.X
	})
}

func getNeighbors(point Point, input [][]string) []Point {
	var neighbors []Point

	for _, diff := range []Point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		neighbor := Point{point.X + diff.X, point.Y + diff.Y}

		if !strings.Contains("EG#", string(input[neighbor.Y][neighbor.X])) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func getTargetPoints(enemies []*Unit, input [][]string) []Point {
	var targets []Point
	seen := make(map[Point]struct{})

	for _, enemy := range enemies {
		if enemy.Health > 0 {
			for _, point := range getNeighbors(enemy.Position, input) {
				if _, ok := seen[point]; !ok {
					seen[point] = struct{}{}
					targets = append(targets, point)
				}
			}
		}
	}

	return targets
}

func pathSearch(start Point, target Point, input [][]string) []Point {
	queue := []Point{start}
	visited := make(map[Point]struct{})
	visited[start] = struct{}{}
	parents := make(map[Point]Point)

	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		if currentPoint == target {
			return createPath(start, target, parents)
		}

		for _, neighbor := range getNeighbors(currentPoint, input) {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = struct{}{}
				parents[neighbor] = currentPoint

				queue = append(queue, neighbor)
			}
		}
	}

	return []Point{}
}

func createPath(start Point, target Point, parents map[Point]Point) []Point {
	path := []Point{target}
	current := target

	for parents[current] != start {
		current = parents[current]
		path = append([]Point{current}, path...)
	}

	return path
}

func contains(point Point, arr []Point) bool {
	for _, item := range arr {
		if item == point {
			return true
		}
	}

	return false
}
