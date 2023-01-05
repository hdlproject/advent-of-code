package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type clList struct {
	next   *clList
	prev   *clList
	number int
}

func newCLList(number int, prev *clList) *clList {
	newList := &clList{
		number: number,
		prev:   prev,
	}

	if prev != nil {
		prev.next = newList
	}

	return newList
}

func day20SolutionA() {
	f, err := os.Open("day_20_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var initSeq []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		valInt, _ := strconv.ParseInt(value, 10, 64)
		initSeq = append(initSeq, int(valInt))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	list := buildCLList(initSeq)

	for index, item := range initSeq {
		decrypt(list, index, item)
	}

	var total int
	startList := list[getMapKey(3507, 0)]
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			total += startList.number
		}

		startList = startList.next
	}

	fmt.Println(total)
}

func buildCLList(initSeq []int) map[string]*clList {
	list := make(map[string]*clList)

	var prevList *clList
	var currectList *clList
	var headList *clList

	for index, item := range initSeq {
		currectList = newCLList(item, prevList)
		if prevList == nil {
			headList = currectList
		}
		prevList = currectList

		list[getMapKey(index, item)] = currectList
	}

	currectList.next = headList
	headList.prev = currectList

	return list
}

func getMapKey(index, item int) string {
	return fmt.Sprintf("%d-%d", index, item)
}

func decrypt(list map[string]*clList, index, item int) {
	currItem := list[getMapKey(index, item)]

	step := int(math.Abs(float64(item))) % (len(list) - 1)
	if step == 0 {
		return
	}

	replacedItem := list[getMapKey(index, item)]
	if item < 0 {
		for i := 0; i <= step; i++ {
			replacedItem = replacedItem.prev
		}
	} else {
		for i := 0; i < step; i++ {
			replacedItem = replacedItem.next
		}
	}

	currItem.prev.next = currItem.next
	currItem.next.prev = currItem.prev

	replacedItem.next.prev = currItem
	currItem.next = replacedItem.next

	replacedItem.next = currItem
	currItem.prev = replacedItem
}

func printList(list map[int]*clList) {
	startList := list[0]
	var finishList *clList

	for startList != finishList {
		fmt.Print(startList.number, ", ")
		startList = startList.next

		if finishList == nil {
			finishList = list[0]
		}
	}
	fmt.Println()
}

func printListBackward(list map[int]*clList) {
	startList := list[0]
	var finishList *clList

	for startList != finishList {
		fmt.Print(startList.number, ", ")
		startList = startList.prev

		if finishList == nil {
			finishList = list[0]
		}
	}
	fmt.Println()
}

func day20SolutionB() {
	f, err := os.Open("day_20_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var initSeq []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		valInt, _ := strconv.ParseInt(value, 10, 64)
		initSeq = append(initSeq, int(valInt))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	list := buildCLList(initSeq)

	for i := 0; i < 10; i++ {
		for index, item := range initSeq {
			decryptType2(list, index, item)
		}
	}

	var total int
	startList := list[getMapKey(3507, 0)]
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			fmt.Println(startList.number)
			total += startList.number
		}

		startList = startList.next
	}

	fmt.Println(total * 811589153)
}

func decryptType2(list map[string]*clList, index, item int) {
	currItem := list[getMapKey(index, item)]

	step := (int(math.Abs(float64(item))) * 811589153) % (len(list) - 1)
	if step == 0 {
		return
	}

	replacedItem := list[getMapKey(index, item)]
	if item < 0 {
		for i := 0; i <= step; i++ {
			replacedItem = replacedItem.prev
		}
	} else {
		for i := 0; i < step; i++ {
			replacedItem = replacedItem.next
		}
	}

	currItem.prev.next = currItem.next
	currItem.next.prev = currItem.prev

	replacedItem.next.prev = currItem
	currItem.next = replacedItem.next

	replacedItem.next = currItem
	currItem.prev = replacedItem
}

//func main() {
//	day20SolutionB()
//}
