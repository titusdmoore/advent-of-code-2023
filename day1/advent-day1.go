package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	// f, err := os.Open("input-short.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var total uint32

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		r := regexp.MustCompile(`\d`)
		matches := r.FindAllString(scanner.Text(), -1)

		if matches != nil {
			var lineTotal uint8
			if len(matches) == 1 {
				matchValue, err := strconv.ParseUint(matches[0], 10, 8)

				if err != nil {
					log.Fatal(err)
				}

				lineTotal = (uint8(matchValue) * 10) + uint8(matchValue)
			} else {
				tens, err := strconv.ParseUint(matches[0], 10, 8)

				if err != nil {
					log.Fatal(err)
				}

				ones, err := strconv.ParseUint(matches[len(matches)-1], 10, 8)

				if err != nil {
					log.Fatal(err)
				}

				lineTotal = (uint8(tens) * 10) + uint8(ones)
			}

			total += uint32(lineTotal)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func part2() {
	// f, err := os.Open("input-short.txt")
	// f, err := os.Open("input.txt")
	// f, err := os.Open("sanity.txt")
	// f, err := os.Open("sanity-part1.txt")
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var total uint32

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		r := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
		matches := r.FindAllString(scanner.Text(), -1)

		if matches != nil {
			var lineTotal uint8
			if len(matches) == 1 {
				matchValue, err := strconv.ParseUint(matches[0], 10, 8)

				if err != nil {
					valFromString, err := StringToInt(matches[0])

					if err != nil {
						log.Fatal(err)
					}

					matchValue = uint64(valFromString)
				}

				lineTotal = (uint8(matchValue) * 10) + uint8(matchValue)
			} else {
				tens, err := strconv.ParseUint(matches[0], 10, 8)

				if err != nil {
					valFromString, err := StringToInt(matches[0])

					if err != nil {
						log.Fatal(err)
					}

					tens = uint64(valFromString)
				}

				ones, err := strconv.ParseUint(matches[len(matches)-1], 10, 8)

				if err != nil {
					valFromString, err := StringToInt(matches[len(matches)-1])

					if err != nil {
						log.Fatal(err)
					}

					ones = uint64(valFromString)
				}

				lineTotal = (uint8(tens) * 10) + uint8(ones)
			}

			total += uint32(lineTotal)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func StringToInt(value string) (int, error) {
	switch value {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return -1, errors.New("Invalid string number value.")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Parameter to specify which part to run. Default 1.")

	flag.Parse()

	switch part {
	case 1:
		part1()
	case 2:
		part2()
	default:
		fmt.Println("Invalid Flag for part")
	}
}
