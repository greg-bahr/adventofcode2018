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
	file, err := ioutil.ReadFile("./day5/input.txt")
	check(err)

	currentString := string(file)
	//currentString = "dabAcCaCBAcCcaDA"

	fmt.Println(len(react(currentString)))

	min := 99999999999999
	for i := 0; i < 26; i++ {
		test := strings.Replace(string(file), string(65+i), "", -1)
		test = strings.Replace(test, string(65+32+i), "", -1)

		if len(react(test)) < min {
			min = len(react(test))
		}
	}

	fmt.Println(min)
}

func react(polymer string) string {
	for i := 0; i < len(polymer)-1; i++ {
		if polymer[i]+32 == polymer[i+1] || polymer[i+1]+32 == polymer[i] {
			polymer = polymer[:i] + polymer[i+2:]

			if i != 0 {
				i -= 2
			} else {
				i--
			}
		}
	}

	return polymer
}
