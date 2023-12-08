package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name  string
	left  *Node
	right *Node
}

func main() {
	f, err := os.Open("day_8/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var instructions []string
	var tree = make(map[string]*Node)

	for scanner.Scan() {
		var line = scanner.Text()

		if len(instructions) == 0 {
			instructions = strings.Split(line, "")
			continue
		}

		if line == "" {
			continue
		}

		var splitLine = strings.Split(line, " = ")
		var nodeName = splitLine[0]
		var children = strings.Split(strings.Replace(strings.Replace(splitLine[1], "(", "", -1), ")", "", -1), ", ")
		var leftChildName = children[0]
		var rightChildName = children[1]

		if _, ok := tree[nodeName]; !ok {
			tree[nodeName] = &Node{name: nodeName}
		}

		if _, ok := tree[leftChildName]; !ok {
			tree[leftChildName] = &Node{name: leftChildName}
		}

		if _, ok := tree[rightChildName]; !ok {
			tree[rightChildName] = &Node{name: rightChildName}
		}

		var mainNode = tree[nodeName]
		var leftChildNode = tree[leftChildName]
		var rightChildNode = tree[rightChildName]

		mainNode.left = leftChildNode
		mainNode.right = rightChildNode
	}

	var rootNode = tree["AAA"]
	var stepsByInstructionsForSingleStartingNode = countStepsByInstructionsForSingleStartingNode(rootNode, instructions, false)
	var stepsByInstructionsForMultipleStartingNodes = countStepsByInstructionsForMultipleStartingNodes(tree, instructions)

	fmt.Printf("Steps by instructions for single starting node: %v \n", stepsByInstructionsForSingleStartingNode)
	fmt.Printf("Steps by instructions for multiple starting nodes: %v \n", stepsByInstructionsForMultipleStartingNodes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countStepsByInstructionsForSingleStartingNode(root *Node, instructions []string, anyEndNodeName bool) int {
	var steps = 0
	var currentInstructionIndex = 0
	currentNode := root

	for {
		if anyEndNodeName && strings.Contains(currentNode.name, "Z") {
			break
		}

		if !anyEndNodeName && currentNode.name == "ZZZ" {
			break
		}

		var currentInstruction = instructions[currentInstructionIndex]

		if currentInstruction == "L" {
			currentNode = currentNode.left
		}

		if currentInstruction == "R" {
			currentNode = currentNode.right
		}

		steps += 1
		currentInstructionIndex += 1

		if currentInstructionIndex == len(instructions) {
			currentInstructionIndex = 0
		}
	}

	return steps
}

func countStepsByInstructionsForMultipleStartingNodes(tree map[string]*Node, instructions []string) int {
	var stepsForEachNode []int

	for key := range tree {
		if strings.Contains(key, "A") {
			var stepsForNode = countStepsByInstructionsForSingleStartingNode(tree[key], instructions, true)
			stepsForEachNode = append(stepsForEachNode, stepsForNode)
		}
	}

	return LCM(stepsForEachNode...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]

	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func tryConvertStrToInt(value string) int {
	if intValue, error := strconv.Atoi(value); error == nil {
		return intValue
	}

	return -1
}
