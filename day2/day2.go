package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day2/input.txt")
	check(err)

	lines := strings.Split(string(file), "\r\n")

	twoCount := 0
	threeCount := 0

	for _, line := range lines {
		if repeats(line, 2) {
			threeCount++
		}
		if repeats(line, 3) {
			twoCount++
		}
	}

	fmt.Println(twoCount * threeCount)

	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			unique, common := diff(lines[i], lines[j])

			if unique == 1 {
				fmt.Println(common)
				return
			}
		}
	}
}

func diff(word1 string, word2 string) (int, string) {
	common := ""
	diff := 0

	for i, l := range word1 {
		if int32(word2[i]) == l {
			common += string(l)
		} else {
			diff++
		}
	}

	return diff, common
}

func repeats(word string, amount int) bool {
	freq := make(map[string]int)

	for _, l := range word {
		freq[string(l)] += 1
	}

	for _, v := range freq {
		if v == amount {
			return true
		}
	}

	return false
}
