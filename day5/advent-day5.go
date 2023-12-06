package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	r := regexp.MustCompile(`(\d+)`)
	var seeds []uint

	for scanner.Scan() {
		if !isWhitespace(scanner.Text()) {
			fmt.Println(scanner.Text())
			switch stage {
			case 0:
				seedsStr := r.FindAllString(scanner.Text(), -1)
				for _, seedStr := range seedsStr {
					seed, err := strconv.ParseUint(seedStr, 10, 64)

					if err != nil {
						log.Fatal(err)
					}

					seeds = append(seeds, uint(seed))
				}
			case 1:
			}

			continue
		}

		stage++
	}

	fmt.Println(seeds)
}
