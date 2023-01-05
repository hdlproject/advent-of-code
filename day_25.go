package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day25SolutionA() {
	f, err := os.Open("day_25_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var decimalVal int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		decimalVal += decipherSNAFU(value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(decimalVal)

	snafuVal := cipherSNAFU(decimalVal)

	fmt.Println(snafuVal)
}

func decipherSNAFU(value string) int64 {
	values := strings.Split(value, "")

	multiplier := int64(1)
	var decimalVal int64
	for i := len(values) - 1; i >= 0; i-- {
		switch values[i] {
		case "-":
			decimalVal -= multiplier
		case "=":
			decimalVal -= 2 * multiplier
		default:
			val, _ := strconv.ParseInt(values[i], 10, 64)
			decimalVal += val * multiplier
		}
		multiplier *= 5
	}

	return decimalVal
}

func cipherSNAFU(value int64) string {
	var strVal string

	mod := int64(5)
	for value > 0 {
		reminder := value % mod

		if reminder == 3 {
			strVal = "=" + strVal
			value += 2
		} else if reminder == 4 {
			strVal = "-" + strVal
			value += 1
		} else {
			reminderStr := strconv.FormatInt(reminder, 10)
			strVal = reminderStr + strVal
		}

		value /= mod
	}

	return strVal
}

func day25SolutionB() {

}

func main() {
	day25SolutionA()
}
