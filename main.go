package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	var letter string
	list, err := readFile()
	var word []rune = []rune(list)
	var underscore []rune
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	underscore = append(underscore)
	fmt.Print("Enter a letter: ")
	fmt.Scanln(&letter)
	for i := 0; i < 10; i++ {
		if string(word[i]) == letter {
			fmt.Print(string(word[i]))
		} else {
			fmt.Print("_")
		}
	}
}

func readFile() (string, error) {
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

	randIndex := rand.Intn(len(list))
	return list[randIndex], nil
}
