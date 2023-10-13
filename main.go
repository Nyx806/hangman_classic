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
	var word []rune
	var underscore []rune
	var list []string

	/* readfile  */
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
	word = []rune(list[randIndex])
	/* end readfile */

	randIndexWord := rand.Intn(len(word))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	for i := 0; i < 10; i++ {
		underscore[randIndexWord] = word[randIndexWord]
		fmt.Println(string(underscore))
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&letter)

		for i := 0; i < len(word); i++ {

			if string(word[i]) == letter {
				underscore[i] = word[i]
			}
		}

		if string(word) == string(underscore) {
			fmt.Println("You win!")
			break
		}
	}
}
