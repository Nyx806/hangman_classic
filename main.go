package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	jeu()
}

func jeu() {
	var remainLifeStr string
	var remainLife int = 9
	var life int = -1
	var letter string
	var underscore []rune
	var index int
	word := readFile()

	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	for i := 0; i < 10; i++ {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		affichage(string(underscore))
		fmt.Print("\n")
		fmt.Print("\n")
		affichage("Enter a letter: ")
		fmt.Println("")
		fmt.Scanln(&letter)
		remainLifeStr = strconv.Itoa(remainLife)

		if remainLife != 1 || remainLife == 1 {
			affichage(" wrong you have" + " " + remainLifeStr + " " + "lives left")
			fmt.Println("")
		}

		for i := 0; i < len(word); i++ {
			if string(word[i]) == letter {
				underscore[i] = word[i]
				index = i
			}

		}

		if string(word[index]) != letter || letter == "" {
			remainLife = 7 - life
			life = life + 1
			lign := life*7 + life

			for i := lign; i < lign+7; i++ {
				fmt.Println(readHangman()[i])
			}
			fmt.Println("")
		}

		if remainLife == -1 {
			fmt.Println("You lose! Try again")
		}

		if string(word) == string(underscore) {
			fmt.Println("You win!")
			break
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

func affichage(n string) {
	for i := 0; i < len(n); i++ {
		fmt.Print(string(n[i]))
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
}
