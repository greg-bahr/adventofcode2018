package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := 635041

	recipes := []int{3, 7}

	elf1 := 0
	elf2 := 1

	for len(recipes) < input*32 {
		sum := strconv.Itoa(recipes[elf1] + recipes[elf2])

		for _, char := range sum {
			num, _ := strconv.Atoi(string(char))

			recipes = append(recipes, num)
		}

		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}

	fmt.Println(recipes[input : input+10])
	fmt.Println(strings.Index(strings.Trim(strings.Replace(fmt.Sprint(recipes), " ", "", -1), "[]"), strconv.Itoa(input)))
}
