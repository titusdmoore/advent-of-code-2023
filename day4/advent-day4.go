package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	// "strconv"
	// "strings"
)

func main() {
	f, err := os.Open("data.txt")
	// f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile(`\d+`)

	total := 0

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ":")
		workingValues := strings.Split(input[1], "|")
		matches := 0

		winningValues := r.FindAllString(workingValues[0], -1)
		ourValues := r.FindAllString(workingValues[1], -1)

		var winningMap map[string]bool
		winningMap = make(map[string]bool)

		for _, value := range winningValues {
			winningMap[value] = true
		}

		for _, value := range ourValues {
			if winningMap[value] {
				matches++
			}
		}

		if matches > 0 {
			total += int(math.Pow(2, float64(matches-1)))
		}
	}

	fmt.Println(total)
}
