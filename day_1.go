package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type maxCalories []int64

func (c maxCalories) Len() int           { return len(c) }
func (c maxCalories) Less(i, j int) bool { return c[i] > c[j] }
func (c maxCalories) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func day1SolutionA() {
	f, err := os.Open("day_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var maxCalorieSum int64
	var calorieSum int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		if value != "" {
			calorie, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			calorieSum += calorie
		} else {
			if calorieSum > maxCalorieSum {
				maxCalorieSum = calorieSum
			}
			calorieSum = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(maxCalorieSum)
}

func day1SolutionB() {
	f, err := os.Open("day_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var calorieSums maxCalories
	var calorieSum int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		if value != "" {
			calorie, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			calorieSum += calorie
		} else {
			calorieSums = append(calorieSums, calorieSum)
			calorieSum = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(calorieSums)

	fmt.Print(calorieSums[0] + calorieSums[1] + calorieSums[2])
}

//func main() {
//	day1Solution()
//}
