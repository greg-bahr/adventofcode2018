package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	Children []Node
	Metadata []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("./day8/input.txt")
	check(err)

	input := strings.Split(string(file), " ")

	part1, _, rootNode := generateTree(0, input)

	fmt.Println(part1)
	fmt.Println(findValue(rootNode))
}

func generateTree(start int, input []string) (sum int, length int, tree Node) {
	sum = 0
	childAmount, _ := strconv.Atoi(input[start])
	metaAmount, _ := strconv.Atoi(input[start+1])
	length = start + 2

	tree = Node{Metadata: []int{}, Children: []Node{}}

	for i := 0; i < childAmount; i++ {
		childSum, childLength, childTree := generateTree(length, input)

		tree.Children = append(tree.Children, childTree)

		sum += childSum
		length = childLength
	}

	for i := 0; i < metaAmount; i++ {
		meta, _ := strconv.Atoi(input[length])

		tree.Metadata = append(tree.Metadata, meta)

		sum += meta
		length++
	}

	return sum, length, tree
}

func findValue(tree Node) int {
	value := 0

	if len(tree.Children) == 0 {
		for _, val := range tree.Metadata {
			value += val
		}

		return value
	}

	for _, meta := range tree.Metadata {
		if meta-1 < len(tree.Children) {
			//fmt.Println(meta, tree.Children[meta])
			value += findValue(tree.Children[meta-1])
		}
	}

	return value
}
