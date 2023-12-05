package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type MapEntry struct {
	destinationRangeStart uint
	sourceRangeStart      uint
	rangeLength           uint
}

func isWhitespace(input string) bool {
	r := regexp.MustCompile(`[\w:\d]+`)

	return !r.MatchString(input)
}

func main() {
	f, err := os.Open("sample.txt")
	// f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	stage := 0

	for scanner.Scan() {
		if !isWhitespace(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}
}
