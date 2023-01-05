package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type listNode struct {
	parentNode *listNode
	items      []*listNode
	value      int
	isBranch   bool
	realValue  string
}

func (l *listNode) addBranch() *listNode {
	newNode := &listNode{
		parentNode: l,
		isBranch:   true,
	}
	l.items = append(l.items, newNode)

	return newNode
}

func (l *listNode) addElement(value int) {
	newNode := &listNode{
		parentNode: l,
		value:      value,
	}
	l.items = append(l.items, newNode)
}

func day13SolutionA() {
	f, err := os.Open("day_13_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nodes, roots []*listNode
	scanner := bufio.NewScanner(f)
	var row int
	var total int
	index := 1
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			index++
			row++
			continue
		}

		values := strings.Split(value, "")

		if row%3 == 0 {
			nodes = []*listNode{{parentNode: nil}, {parentNode: nil}}
			roots = nodes
		}

		node := nodes[row%3]
		var numberStr string
		for i := 0; i < len(values); i++ {
			if numberStr != "" && !isNumber(values[i]) {
				element, _ := strconv.ParseInt(numberStr, 10, 64)
				node.addElement(int(element))
				numberStr = ""
			}

			switch values[i] {
			case "[":
				node = node.addBranch()
			case "]":
				node = node.parentNode
			case ",":
			case " ":
			default:
				numberStr += values[i]
			}
		}

		if row%3 == 1 {
			isValid, _ := compare(roots[0].items[0], roots[1].items[0])
			if isValid {
				total += index
			}
		}

		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

type sortedListNodes []*listNode

func (s sortedListNodes) Len() int { return len(s) }
func (s sortedListNodes) Less(i, j int) bool {
	temp1 := s[i].items[0]
	temp2 := s[j].items[0]
	isValid, _ := compare(temp1, temp2)
	return isValid
}
func (s sortedListNodes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func day13SolutionB() {
	f, err := os.Open("day_13_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var nodes sortedListNodes
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			continue
		}

		values := strings.Split(value, "")

		newNode := &listNode{realValue: value}
		nodes = append(nodes, newNode)
		var numberStr string
		for i := 0; i < len(values); i++ {
			if numberStr != "" && !isNumber(values[i]) {
				element, _ := strconv.ParseInt(numberStr, 10, 64)
				newNode.addElement(int(element))
				numberStr = ""
			}

			switch values[i] {
			case "[":
				newNode = newNode.addBranch()
			case "]":
				newNode = newNode.parentNode
			case ",":
			case " ":
			default:
				numberStr += values[i]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(nodes)

	key := 1
	for index, node := range nodes {
		if node.realValue == "[[2]]" || node.realValue == "[[6]]" {
			key = key * (index + 1)
		}
	}

	fmt.Println(key)
}

func isNumber(char string) bool {
	return char != "[" &&
		char != "]" &&
		char != "," &&
		char != " "
}

func printTree(node *listNode, indentation int) {
	for _, item := range node.items {
		if !item.isBranch {
			for i := 0; i < indentation; i++ {
				fmt.Print("-")
			}
			fmt.Println(item.value)
		} else {
			for i := 0; i < indentation; i++ {
				fmt.Print("-")
			}
			fmt.Println()
		}
		printTree(item, indentation+1)
	}
}

func compare(node1 *listNode, node2 *listNode) (isValid bool, isContinue bool) {
	if node1.isBranch == node2.isBranch {
		if !node1.isBranch {
			if node1.value > node2.value {
				return false, false
			} else if node1.value == node2.value {
				return true, true
			} else {
				return true, false
			}
		} else {
			for index, item := range node1.items {
				if index < len(node2.items) {
					isValid, isContinue = compare(item, node2.items[index])
					if !isContinue {
						return isValid, false
					}
				}
			}

			isValid = len(node1.items) <= len(node2.items)
			isContinue = len(node1.items) == len(node2.items)

			return isValid, isContinue
		}
	} else {
		if !node1.isBranch {
			newNode := wrapNode(node1)
			isValid, isContinue = compare(newNode, node2)
			return isValid, isContinue
		} else {
			newNode := wrapNode(node2)
			isValid, isContinue = compare(node1, newNode)
			return isValid, isContinue
		}
	}
}

func wrapNode(node *listNode) *listNode {
	newNode := &listNode{
		items:    []*listNode{node},
		isBranch: true,
	}
	node = newNode

	return node
}

//func main() {
//	day13SolutionB()
//}
