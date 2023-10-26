package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
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
	var bold string = "\033[1m"

	affichage(green + bold + "Welcome to Hangman game ! \n" + reset)
	fmt.Println("")
	affichage(green + "Enter your username : " + reset)
	fmt.Scanln(&user)
	fmt.Println("")
	affichage(green + bold + "welcome " + user + " ! \n" + reset)
	fmt.Println("")
	affichage(green + "you have the choise betwen 3 dificulty : \n" + reset)
	fmt.Println("")
	affichage(green + "easy \n" + reset)
	affichage(green + "medium \n" + reset)
	affichage(green + "hard \n" + reset)
	fmt.Println("")
	affichage(green + "please make your choise : " + reset)
	fmt.Scanln(&choise)

	for choise != "easy" && choise != "medium" && choise != "hard" {
		fmt.Println("")
		affichage(green + "incorrect input please retry : \n" + reset)
		fmt.Println("")
		affichage(green + "please make your choise : " + reset)
		fmt.Scanln(&choise)
		fmt.Println("")
	}

	var remainLifeStr string
	var remainLife int = 10
	var life int = -1
	var letter string
	var underscore []rune
	var index int
	var count int
	var interupEasy int = 0
	var interupMedium int = 0
	var interupHard int = 0

	word := readFile(choise)

	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	for i := 0; i < 10; i++ {
		count++
		fmt.Println(green + "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" + reset)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		affichage(green + string(underscore) + reset)
		fmt.Print("\n")
		fmt.Print("\n")
		affichage(green + "Enter a letter : " + reset)
		fmt.Scanln(&letter)
		fmt.Println("")
		fmt.Println(remainLife)

		if letter == "indice" {
			var indice int
			if choise == "easy" && interupEasy < 2 && underscore[indice] != word[indice] {
				indice = rand.Intn(len(word))
				underscore[indice] = word[indice]
				interupEasy++
			} else if choise == "medium" && interupMedium < 3 && underscore[indice] != word[indice] {
				indice = rand.Intn(len(word))
				underscore[indice] = word[indice]
				interupMedium++
			} else if choise == "hard" && interupHard < 4 && underscore[indice] != word[indice] {
				indice = rand.Intn(len(word))
				underscore[indice] = word[indice]
				interupHard++
			} else {
				affichage(green + "you have use all your indice ! \n" + reset)
				remainLife = remainLife + 1
			}
		}

		if letter == "menu" {
			var desire string
			affichage(green + "are you sure you want to go back to the main menu ? [y/n] : " + reset)
			fmt.Scanln(&desire)
			if desire == "y" {
				affichage(green + "ok let's go back to the main menu ! \n" + reset)
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
				jeu()
			} else if desire == "n" {
				affichage(green + "ok let's continue ! \n" + reset)
			}
		}

		for i := 0; i < len(word); i++ {
			if string(word[i]) == letter {
				underscore[i] = word[i]
				index = i
			}

		}

		if string(word[index]) != letter || letter == "" {
			life = life + 1
			lign := life*7 + life
			remainLife--
			remainLifeStr = strconv.Itoa(remainLife)
			affichage(green + "wrong you have" + " " + remainLifeStr + " " + "lives left" + reset)
			fmt.Println("")

			for i := lign; i < lign+7; i++ {
				fmt.Println(green + readHangman()[i] + reset)
			}
			fmt.Println("")
		}

		if word[index] == rune(letter[0]) && remainLife != 10 {
			affichage(green + " you have" + " " + remainLifeStr + " " + "lives left" + reset)
			fmt.Println("")
		}

		if remainLife == -1 {
			fmt.Println(green + "You lose! Try again" + reset)
			affichage(green + "The word was : " + string(word) + reset)
		}

		if string(word) == string(underscore) {
			fmt.Println(green + "You win! \n" + reset)
			affichage(green + "The word was : " + string(word) + reset)
			break
		}
	}

	if count == 10 {
		affichage(green + "you have use your 10 atempts, you lose ! \n" + reset)
		affichage(green + "The word was : " + string(word) + reset)
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
