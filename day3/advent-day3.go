package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type PNumberPosition struct {
	start int
	end   int
}

type PotentialPartNumber struct {
	value    uint64
	position PNumberPosition
}

func isValidPartNumber(line int, number PotentialPartNumber, input [][]byte) bool {
	out := false
	r := regexp.MustCompile(`[\*\$\=\#\%\/\&\+\-\@\\]`)

	// Check above
	if line > 0 {
		workingLine := input[line-1]
		start := number.position.start
		end := number.position.end

		if start > 0 {
			start--
		}

		bytesAbove := workingLine[start : end+1]

		for _, singleByte := range bytesAbove {
			var bytes []byte
			bytes = append(bytes, singleByte)
			if r.Match(bytes) {
				out = true
			}
		}
	}

	// Check left and right
	if number.position.start > 0 {
		workingLine := input[line]
		singleByte := workingLine[number.position.start-1]

		var bytes []byte
		bytes = append(bytes, singleByte)
		if r.Match(bytes) {
			out = true
		}
	}

	if number.position.end < len(input[line]) {
		workingLine := input[line]
		singleByte := workingLine[number.position.end]

		var bytes []byte
		bytes = append(bytes, singleByte)
		if r.Match(bytes) {
			out = true
		}
	}

	// Check below
	if line+1 < len(input) {
		workingLine := input[line+1]
		start := number.position.start
		end := number.position.end

		if start > 0 {
			start--
		}

		bytesBelow := workingLine[start : end+1]

		// if number.value == 720 {
		// 	fmt.Printf("Line: %d, Start: %d, End: %d, Bytes: %s\n", line, start, end, bytesBelow)
		// 	fmt.Printf("%s\n", input[line])
		// }

		for _, singleByte := range bytesBelow {
			var bytes []byte
			bytes = append(bytes, singleByte)
			if r.Match(bytes) {
				out = true
			}
		}
	}

	return out
}

func main() {
	f, err := os.Open("data.txt")
	// f, err := os.Open("sample.txt")
	// f, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var reference [][]byte
	var potentialPartNumbers [][]PotentialPartNumber
	r := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		reference = append(reference, []byte(scanner.Text()))
		var rowPotentialParts []PotentialPartNumber

		matches := r.FindAllString(scanner.Text(), -1)
		indexes := r.FindAllStringIndex(scanner.Text(), -1)

		for i := range matches {
			val, err := strconv.ParseUint(matches[i], 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			rowPotentialParts = append(rowPotentialParts, PotentialPartNumber{
				value: val,
				position: PNumberPosition{
					start: indexes[i][0],
					end:   indexes[i][1],
				},
			})
		}

		potentialPartNumbers = append(potentialPartNumbers, rowPotentialParts)
	}

	var total uint

	for i, line := range reference {
		fmt.Printf("%d: %s\n", i, line)
	}

	fmt.Println(total)
}
