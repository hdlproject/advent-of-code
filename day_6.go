package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type queue []string

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func day6SolutionA() {
	f, err := os.Open("day_6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, "")

		var queue queue
		for index, char := range values {
			queue.Enqueue(char)
			if len(queue) > 4 {
				queue.Dequeue()
			}

			uniqueMap := make(map[string]struct{})
			for _, item := range queue {
				uniqueMap[item] = struct{}{}
			}

			if len(uniqueMap) == 4 {
				fmt.Println(index + 1)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day6SolutionB() {
	f, err := os.Open("day_6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, "")

		var queue queue
		for index, char := range values {
			queue.Enqueue(char)
			if len(queue) > 14 {
				queue.Dequeue()
			}

			uniqueMap := make(map[string]struct{})
			for _, item := range queue {
				uniqueMap[item] = struct{}{}
			}

			if len(uniqueMap) == 14 {
				fmt.Println(index + 1)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	day6SolutionB()
//}
