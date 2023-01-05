package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	elevationStr string
	col          int
	row          int

	elevation int
	isVisited bool
	distance  int

	top      *node
	bottom   *node
	right    *node
	left     *node
	prevNode *node
}

func newNode(colIndex, rowIndex int, heightMap [][]*node, isAdjacent func(int, int, int, int, [][]*node) bool) *node {
	var top *node
	topIndex := colIndex - 1
	if topIndex >= 0 &&
		isAdjacent(colIndex, rowIndex, topIndex, rowIndex, heightMap) {

		top = heightMap[topIndex][rowIndex]
	}

	var bottom *node
	bottomIndex := colIndex + 1
	if bottomIndex < len(heightMap) &&
		isAdjacent(colIndex, rowIndex, bottomIndex, rowIndex, heightMap) {

		bottom = heightMap[bottomIndex][rowIndex]
	}

	var right *node
	rightIndex := rowIndex + 1
	if rightIndex < len(heightMap[0]) &&
		isAdjacent(colIndex, rowIndex, colIndex, rightIndex, heightMap) {

		right = heightMap[colIndex][rightIndex]
	}

	var left *node
	leftIndex := rowIndex - 1
	if leftIndex >= 0 &&
		isAdjacent(colIndex, rowIndex, colIndex, leftIndex, heightMap) {

		left = heightMap[colIndex][leftIndex]
	}

	heightMap[colIndex][rowIndex].top = top
	heightMap[colIndex][rowIndex].bottom = bottom
	heightMap[colIndex][rowIndex].right = right
	heightMap[colIndex][rowIndex].left = left

	return heightMap[colIndex][rowIndex]
}

func isAdjacentType1(currCol, currRow, nextCol, nextRow int, heightMap [][]*node) bool {
	return heightMap[nextCol][nextRow].elevation-heightMap[currCol][currRow].elevation <= 1
}

func isAdjacentType2(currCol, currRow, nextCol, nextRow int, heightMap [][]*node) bool {
	return heightMap[currCol][currRow].elevation-heightMap[nextCol][nextRow].elevation <= 1
}

type queueNode []*node

func (q *queueNode) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queueNode) Enqueue(node *node) {
	*q = append(*q, node)
}

func (q *queueNode) Dequeue() (*node, bool) {
	if q.IsEmpty() {
		return nil, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func day12SolutionA() {
	f, err := os.Open("day_12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var heightMap [][]*node
	var startCol, startRow int
	var finishCol, finishRow int
	var currentCol int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, "")

		var heightRow []*node
		for rowIndex, item := range values {
			if item == "S" {
				startCol = currentCol
				startRow = rowIndex
				item = "a"
			} else if item == "E" {
				finishCol = currentCol
				finishRow = rowIndex
				item = "z"
			}

			itemInt := getElevationInt(item)

			heightRow = append(heightRow, &node{
				elevationStr: item,
				elevation:    itemInt,
				col:          currentCol,
				row:          rowIndex,
			})
		}
		heightMap = append(heightMap, heightRow)
		currentCol++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for colIndex, row := range heightMap {
		for rowIndex := range row {
			heightMap[colIndex][rowIndex] = newNode(colIndex, rowIndex, heightMap, isAdjacentType1)
		}
	}

	startNode := heightMap[startCol][startRow]
	finishNode := heightMap[finishCol][finishRow]

	shortestDistance := bfs(startNode, getIsFinishedType1(finishNode))

	fmt.Println(shortestDistance)
}

func day12SolutionB() {
	f, err := os.Open("day_12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var heightMap [][]*node
	var finishCol, finishRow int
	var currentCol int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, "")

		var heightRow []*node
		for rowIndex, item := range values {
			if item == "S" {
				item = "a"
			} else if item == "E" {
				finishCol = currentCol
				finishRow = rowIndex
				item = "z"
			}

			itemInt := getElevationInt(item)

			heightRow = append(heightRow, &node{
				elevationStr: item,
				elevation:    itemInt,
				col:          currentCol,
				row:          rowIndex,
			})
		}
		heightMap = append(heightMap, heightRow)
		currentCol++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for colIndex, row := range heightMap {
		for rowIndex := range row {
			heightMap[colIndex][rowIndex] = newNode(colIndex, rowIndex, heightMap, isAdjacentType2)
		}
	}

	finishNode := heightMap[finishCol][finishRow]

	shortestDistance := bfs(finishNode, getIsFinishedType2())

	fmt.Println(shortestDistance)
}

func getElevationInt(str string) int {
	strRune := []rune(str)[0]
	return int(strRune) - 96
}

func bfs(startNode *node, isFinished func(*node) bool) int {
	var queue queueNode
	startNode.isVisited = true
	queue.Enqueue(startNode)

	var currNode *node
	for !queue.IsEmpty() {
		currNode, _ = queue.Dequeue()

		if isFinished(currNode) {
			return currNode.distance
		}

		if currNode.top != nil && !currNode.top.isVisited {
			currNode.top.isVisited = true
			currNode.top.distance = currNode.distance + 1
			currNode.top.prevNode = currNode
			queue.Enqueue(currNode.top)
		}
		if currNode.bottom != nil && !currNode.bottom.isVisited {
			currNode.bottom.isVisited = true
			currNode.bottom.distance = currNode.distance + 1
			currNode.bottom.prevNode = currNode
			queue.Enqueue(currNode.bottom)
		}
		if currNode.right != nil && !currNode.right.isVisited {
			currNode.right.isVisited = true
			currNode.right.distance = currNode.distance + 1
			currNode.right.prevNode = currNode
			queue.Enqueue(currNode.right)
		}
		if currNode.left != nil && !currNode.left.isVisited {
			currNode.left.isVisited = true
			currNode.left.distance = currNode.distance + 1
			currNode.left.prevNode = currNode
			queue.Enqueue(currNode.left)
		}
	}

	return -1
}

func getIsFinishedType1(finishNode *node) func(*node) bool {
	return func(currNode *node) bool {
		return currNode == finishNode
	}
}

func getIsFinishedType2() func(*node) bool {
	return func(currNode *node) bool {
		return currNode.elevationStr == "a"
	}
}

//func main() {
//	day12SolutionB()
//}
