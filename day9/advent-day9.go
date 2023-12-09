package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringArrToNumArr(strArr []string) []int {
	var numArr []int

	for _, str := range strArr {
		num, _ := strconv.Atoi(str)
		numArr = append(numArr, num)
	}

	return numArr
}

func findNextDigit(numbers []int) int {
	var data [][]int
	data = append(data, numbers)

	findNextDigitHelper(&data, 0)

	// Minus 2 because -1 to prevent out of range, and -1 because no need to run on full 0 arr
	for i := len(data) - 2; i >= 0; i-- {
		newVal := data[i+1][len(data[i+1])-1] + data[i][len(data[i])-1]
		data[i] = append(data[i], newVal)
	}

	return data[0][len(data[0])-1]
}

func findNextDigitHelper(data *[][]int, run int) {
	if arrIsZeroed((*data)[len(*data)-1]) {
		return
	}

	var rowDiff []int
	numbers := (*data)[run]

	for i, number := range numbers {
		if i > 0 {
			rowDiff = append(rowDiff, number-numbers[i-1])
		}
	}

	*data = append(*data, rowDiff)
	findNextDigitHelper(data, run+1)
}

func arrIsZeroed(input []int) bool {
	out := true

	for _, val := range input {
		out = out && val == 0
	}

	return out
}

func main() {
	f, err := os.Open("data.txt")
	// f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		stringNumbers := strings.Split(line, " ")
		numbers := stringArrToNumArr(stringNumbers)

		newLast := findNextDigit(numbers)
		total += newLast
		fmt.Println(newLast)
	}

	fmt.Println("total", total)
}
