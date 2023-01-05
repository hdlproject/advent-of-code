package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	parent *dir
	dirs   map[string]*dir
	files  map[string]int64
}

func createDir(parent *dir) *dir {
	newDir := &dir{
		dirs:  make(map[string]*dir),
		files: make(map[string]int64),
	}
	newDir.parent = parent

	return newDir
}

var expectedTotal int64
var neededSpace int64

func day7SolutionA() {
	f, err := os.Open("day_7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dir := createDir(nil)
	root := dir
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		dir = buildStruct(value, dir)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	calculateTotal(root)

	fmt.Println(expectedTotal)
}

func day7SolutionB() {
	f, err := os.Open("day_7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dir := createDir(nil)
	root := dir
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		dir = buildStruct(value, dir)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := calculateTotal(root)
	neededSpace = 30000000 - (70000000 - total)

	expectedTotal = math.MaxInt64
	findMinimum(root)

	fmt.Println(expectedTotal)
}

func buildStruct(value string, dir *dir) *dir {
	values := strings.Split(value, " ")

	switch values[0] {
	case "$":
		switch values[1] {
		case "cd":
			switch values[2] {
			case ".":
				// do nothing
			case "..":
				dir = dir.parent
			default:
				if _, ok := dir.dirs[values[2]]; !ok {
					dir.dirs[values[2]] = createDir(dir)
				}
				dir = dir.dirs[values[2]]
			}
		case "ls":
			// do nothing
		}
	case "dir":
		if _, ok := dir.dirs[values[1]]; !ok {
			dir.dirs[values[1]] = createDir(dir)
		}
	default:
		amount, _ := strconv.ParseInt(values[0], 10, 64)
		dir.files[values[1]] = amount
	}

	return dir
}

func calculateTotal(dir *dir) int64 {
	var total int64
	for _, file := range dir.files {
		total += file
	}

	for _, child := range dir.dirs {
		total += calculateTotal(child)
	}

	if total < 100000 {
		expectedTotal += total
	}

	return total
}

func findMinimum(dir *dir) int64 {
	var total int64
	for _, file := range dir.files {
		total += file
	}

	for _, child := range dir.dirs {
		total += findMinimum(child)
	}

	if total >= neededSpace {
		expectedTotal = int64(math.Min(float64(expectedTotal), float64(total)))
	}

	return total
}

//func main() {
//	day7SolutionA()
//}
