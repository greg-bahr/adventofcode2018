package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day4/input.txt")
	check(err)

	lines := strings.Split(string(file), "\r\n")
	sort.Strings(lines)

	re := regexp.MustCompile("[0-9]+")

	guards := make(map[int][]int)

	currentGuard := -1
	for currentLocation := 0; currentLocation < len(lines); currentLocation++ {
		line := lines[currentLocation]
		nums := re.FindAllString(line, -1)

		minute, _ := strconv.Atoi(nums[4])

		if strings.Contains(line, "Guard") {
			id, _ := strconv.Atoi(nums[5])

			currentGuard = id

			if _, ok := guards[currentGuard]; !ok {
				guards[currentGuard] = make([]int, 60)
			}
		} else if strings.Contains(line, "falls") {
			currentLocation++

			wakeMinute, _ := strconv.Atoi(re.FindAllString(lines[currentLocation], -1)[4])

			for i := minute; i < wakeMinute; i++ {
				guards[currentGuard][i] += 1
			}
		}
	}

	laziestGuard := -1
	mostAsleep := -1

	for id, minutes := range guards {
		sum := 0

		for _, num := range minutes {
			sum += num
		}

		if sum > mostAsleep {
			laziestGuard = id
			mostAsleep = sum
		}
	}

	bestMinute := -1
	timesAsleep := -1

	for i, times := range guards[laziestGuard] {
		if times > timesAsleep {
			timesAsleep = times
			bestMinute = i
		}
	}

	fmt.Println(bestMinute * laziestGuard)

	maxAsleep := -1
	worstGuard := -1
	minuteNumber := -1

	for id, minutes := range guards {
		for i, times := range minutes {
			if times > maxAsleep {
				maxAsleep = times
				worstGuard = id
				minuteNumber = i
			}
		}
	}

	fmt.Println(worstGuard * minuteNumber)
}
