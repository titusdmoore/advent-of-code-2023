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
	f, err := os.Open("input.txt")
	// f, err := os.Open("sanity.txt")
	// f, err := os.Open("sanity-part1.txt")
	// f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var total uint32

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineVal := parseInput(scanner.Text())
		total += uint32(lineVal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func parseInput(input string) uint {
	regex := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	var reversedMatches []uint

	lastVal := stringParseInput(input, 0, &reversedMatches)

	if regex.MatchString(lastVal) {
		val, err := strconv.ParseUint(lastVal, 10, 8)

		if err != nil {
			// We are a string value
			val, err := StringToInt(lastVal)

			if err != nil {
				log.Fatal(err)
			}

			reversedMatches = append(reversedMatches, uint(val))
		} else {
			// we were a number value, no need to keep prev
			reversedMatches = append(reversedMatches, uint(val))
		}
	}

	return (reversedMatches[len(reversedMatches)-1] * 10) + reversedMatches[0]
}

func stringParseInput(input string, position int, reversedMatches *[]uint) string {
	if position == len(input) {
		return ""
	}

	prev := stringParseInput(input, position+1, reversedMatches)

	regex := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)$`)

	// So prev is a valid value, but we are not sure if its a number or a string value of number
	if regex.MatchString(prev) {
		fmt.Println("here", prev, input)
		val, err := strconv.ParseUint(prev, 10, 8)

		if err != nil {
			// We are a string value
			val, err := StringToInt(prev)

			if err != nil {
				log.Fatal(err)
			}

			*reversedMatches = append(*reversedMatches, uint(val))
			if matched, err := regexp.MatchString(`\d`, string(input[position-1])); err == nil && matched {
				fmt.Println("Here")
				prev = ""
			} else {
				fmt.Println("Here 2")
				prev = string(prev[0])
			}
		} else {
			// we were a number value, no need to keep prev
			*reversedMatches = append(*reversedMatches, uint(val))
			fmt.Println(reversedMatches)
			prev = ""
		}
	}

	return string(input[position]) + prev
}

// Move to main function
func handleValue(value string, reversedMatches []uint) bool {
	regex := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

	if regex.MatchString(value) {
		val, err := strconv.ParseUint(value, 10, 8)

		if err != nil {
			strVal, err := StringToInt(value)

			if err != nil {
				log.Fatal(err)
			}

			reversedMatches = append(reversedMatches, uint(strVal))
			return true
		}

		reversedMatches = append(reversedMatches, uint(val))
		return true
	}

	return false
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
