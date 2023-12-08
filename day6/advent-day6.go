package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// f, err := os.Open("sample.txt")
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	row := 0
	var times []string
	var distances []string
	totalWiggleRoom := 1

	for scanner.Scan() {
		r := regexp.MustCompile(`\d+`)

		switch row {
		case 0:
			times = r.FindAllString(scanner.Text(), -1)
		case 1:
			distances = r.FindAllString(scanner.Text(), -1)
		}

		row++
	}

	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])

		if err != nil {
			log.Fatal(err)
		}

		distance, err := strconv.Atoi(distances[i])

		if err != nil {
			log.Fatal(err)
		}

		wins := 0

		// Start at 1 because we can't go anywhere with nothing, end 1 before max, because then there is no time
		for j := 1; j < time-1; j++ {
			distanceTraveled := j * (time - j)

			if distanceTraveled > distance {
				wins++
			}
		}

		totalWiggleRoom *= wins
	}

	fmt.Println(totalWiggleRoom)
}
