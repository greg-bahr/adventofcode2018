package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Worker struct {
	Time int
	Step string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day7/input.txt")
	check(err)

	lines := strings.Split(string(file), "\r\n")

	steps := make(map[string][]string)

	for _, line := range lines {
		arr := strings.Split(line, " ")

		step := arr[7]
		prereq := arr[1]

		steps[step] = append(steps[step], prereq)
		steps[prereq] = steps[prereq]
	}

	order := ""
	ran := make(map[string]bool)
	queue := []string{}

	for step, prereqs := range steps {
		if len(prereqs) == 0 {
			queue = append(queue, step)
		}
	}
	sort.Strings(queue)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		ran[current] = true

		for step, prereqs := range steps {
			if !ran[step] && !in(step, queue) && canRun(prereqs, ran) {
				queue = append(queue, step)
			}
		}

		order += current
		sort.Strings(queue)
	}

	fmt.Println(order)

	second := 0
	workers := [5]Worker{}
	ran = make(map[string]bool)
	queue = []string{}

	for step, prereqs := range steps {
		if len(prereqs) == 0 {
			queue = append(queue, step)
		}
	}
	sort.Strings(queue)

	for len(queue) > 0 || isWorking(workers) {
		for i := 0; i < len(workers); i++ {
			if workers[i].Time > 0 {
				workers[i].Time--
			}

			if workers[i].Time == 0 {
				ran[workers[i].Step] = true
				workers[i].Step = "."

				for step, prereqs := range steps {
					if !ran[step] && !in(step, queue) && !isRunning(step, workers) && canRun(prereqs, ran) {
						queue = append(queue, step)
					}
				}
			}
		}

		for i := 0; i < len(workers); i++ {
			if workers[i].Time == 0 && len(queue) > 0 {
				workers[i].Step = queue[0]
				queue = queue[1:]

				time := workers[i].Step[0] - 4

				workers[i].Time = int(time)
			}
		}

		sort.Strings(queue)
		second++
	}

	fmt.Println(second - 1)
}

func isRunning(step string, workers [5]Worker) bool {
	for _, worker := range workers {
		if worker.Step == step {
			return true
		}
	}

	return false
}

func isWorking(workers [5]Worker) bool {
	for _, worker := range workers {
		if worker.Time > 0 {
			return true
		}
	}

	return false
}

func in(ch string, arr []string) bool {
	for _, item := range arr {
		if ch == item {
			return true
		}
	}

	return false
}

func canRun(prereqs []string, ran map[string]bool) bool {
	for _, prereq := range prereqs {
		if !ran[prereq] {
			return false
		}
	}

	return true
}
