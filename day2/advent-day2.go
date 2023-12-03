package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// regex := regexp.MustCompile(`Game (?P<gameNumber>\d): (?P<colors>\d+ ?(red|green|blue)(?:[, ;]?)+)+`)
	regex := regexp.MustCompile(`(\d ?(red|green|blue)?)`)

	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// matches := regex.FindAllString(scanner.Text(), -1)
		//
		// fmt.Println(matches)
		fmt.Println(scanner.Text(), regex.FindAllStringSubmatch(scanner.Text(), -1))
		fmt.Println(regex.SubexpNames())
	}
}
