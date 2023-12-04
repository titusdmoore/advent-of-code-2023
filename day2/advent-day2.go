package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidColor(color string) bool {
	colorValues := strings.Split(strings.TrimSpace(color), " ")

	val, err := strconv.ParseInt(colorValues[0], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	switch colorValues[1] {
	case "red":
		return val <= 12
	case "green":
		return val <= 13
	case "blue":
		return val <= 14
	}

	return false
}

func main() {
	// f, err := os.Open("sample.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		splitWithGame := strings.Split(scanner.Text(), ":")
		r := regexp.MustCompile(`\d+`)
		gameNumber := r.FindString(splitWithGame[0])

		rounds := strings.Split(splitWithGame[1], ";")
		validGame := true

		for _, round := range rounds {
			colors := strings.Split(round, ",")

			for _, color := range colors {
				validGame = validGame && isValidColor(color)
			}
		}

		if validGame {
			val, err := strconv.ParseInt(gameNumber, 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			total += int(val)
		}
	}

	fmt.Println(total)
}
