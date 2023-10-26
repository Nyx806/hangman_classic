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
	var choise string
	var user string
	var green string = "\033[32m"
	var reset string = "\033[0m"

	affichage(green + "Enter your username : " + reset)
	fmt.Scanln(&user)

	affichage(green + "Welcome to Hangman game " + user + " ! \n" + reset)
	affichage(green + "you have the choise betwen 3 dificulty : \n" + reset)
	affichage(green + "easy \n" + reset)
	affichage(green + "medium \n" + reset)
	affichage(green + "hard \n" + reset)
	affichage(green + "please make your choise : " + reset)
	fmt.Scanln(&choise)

	var remainLifeStr string
	var remainLife int = 9
	var life int = -1
	var letter string
	var underscore []rune
	var index int

	word := readFile(choise)

	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	for i := 0; i < 10; i++ {
		fmt.Println(green + "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" + reset)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		affichage(green + string(underscore) + reset)
		fmt.Print("\n")
		fmt.Print("\n")
		affichage(green + "Enter a letter: " + reset)
		fmt.Scanln(&letter)
		remainLifeStr = strconv.Itoa(remainLife)
		fmt.Println("")

		if remainLife != 1 || remainLife == 1 {
			affichage(green + " you have" + " " + remainLifeStr + " " + "lives left" + reset)
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
				fmt.Println(green + readHangman()[i] + reset)
			}
			fmt.Println("")
		}

		if remainLife == -1 {
			fmt.Println(green + "You lose! Try again" + reset)
		}

		if string(word) == string(underscore) {
			fmt.Println(green + "You win!" + reset)
			break
		}
	}
}

func readFile(r string) []rune {

	var word []rune
	var list []string
	var file *os.File
	var err error
	if r == "easy" {
		file, err = os.Open("facile.txt")
	} else if r == "medium" {
		file, err = os.Open("moyen.txt")
	} else if r == "hard" {
		file, err = os.Open("difficile.txt")
	}

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
