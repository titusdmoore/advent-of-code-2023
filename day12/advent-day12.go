package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "regexp"
	"strings"
)

func main() {
	f, err := os.Open("sample.txt")
	// f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")

		fmt.Println(values)
	}
}
