package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type valve struct {
	connectTo map[string]*valve
	flowRate  int
}

func newValve(flowRate int, connectTo []string) *valve {
	connectToMap := make(map[string]*valve)
	for _, anotherValve := range connectTo {
		connectToMap[anotherValve] = &valve{}
	}

	return &valve{
		connectTo: connectToMap,
		flowRate:  flowRate,
	}
}

func day16SolutionA() {
	f, err := os.Open("day_16_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	valveMap := make(map[string]*valve)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		id := values[1]
		flowRate, _ := strconv.ParseInt(strings.TrimPrefix(strings.TrimSuffix(values[4], ";"), "rate="), 10, 64)
		var connectTo []string
		for _, val := range values[9:] {
			connectTo = append(connectTo, strings.TrimSuffix(val, ","))
		}

		valveMap[id] = newValve(int(flowRate), connectTo)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	buildGraph(valveMap)

	findMaxWeight(valveMap["AA"])
}

func buildGraph(valveMap map[string]*valve) {
	for _, currValve := range valveMap {
		for index := range currValve.connectTo {
			currValve.connectTo[index] = valveMap[index]
		}
	}
}

func findMaxWeight(valve *valve) (string, int) {
	var maxWeight int
	var maxIndex string
	for id, anotherValve := range valve.connectTo {
		weight := calculateWeight(valve, 0, anotherValve)
		if weight > maxWeight {
			fmt.Println(maxIndex)
			maxWeight = weight
			maxIndex = id
		}
	}

	return maxIndex, maxWeight
}

func calculateWeight(src *valve, step int, currValve *valve) int {
	var weight int
	for _, anotherValve := range currValve.connectTo {
		if anotherValve == src {
			weight += 0
			continue
		}

		weight += weight + (anotherValve.flowRate * (30 - step)) + calculateWeight(src, step+1, anotherValve)
	}

	return weight
}

func day16SolutionB() {

}

//func main() {
//	day16SolutionA()
//}
