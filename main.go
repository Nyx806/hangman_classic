package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	list, err := readFile()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() ([]string, error) {
	var list []string
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list, nil
}
