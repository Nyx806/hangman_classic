package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	jeu()
}

func jeu() {
	var life int = -1
	var letter string
	var underscore []rune
	//var count int
	var index int
	word := readFile()

	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	for i := 0; i < 10; i++ {
		fmt.Println(string(underscore) + " ")
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&letter)

		for i := 0; i < len(word); i++ {
			if string(word[i]) == letter {
				underscore[i] = word[i]
				index = i
			}

		}

		if string(word[index]) != letter {
			fmt.Println(life)
			life = life + 1
			lign := life*7 + life

			for i := lign; i < lign+7; i++ {
				fmt.Println(readHangman()[i])
			}
		}

		if string(word) == string(underscore) {
			fmt.Println(string(underscore))
			fmt.Println("You win!")
		}
	}
}

func readFile() []rune {

	var word []rune
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
	word = []rune(list[randIndex])
	return word

}

func readHangman() []string {
	var hangman []string

	file, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hangman = append(hangman, scanner.Text())
	}

	// Check for errors

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return hangman
}

func draw(min int, max int) []string {
	tab := readHangman()
	var newtab []string

	for i := min; i <= max; i++ {
		newtab = append(newtab, tab[i])
	}
	return newtab
}
