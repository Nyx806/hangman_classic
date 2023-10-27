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
	var lang string
	var green string = "\033[32m"
	var reset string = "\033[0m"
	var bold string = "\033[1m"

	affichage(green + "select your language [fr] [en] : " + reset)
	fmt.Scanln(&lang)

	for lang != "fr" && lang != "en" {
		affichage(green + "incorrect answer please select your language [fr] [en] : " + reset)
		fmt.Scanln(&lang)
	}

	fmt.Println("")
	selectLanguage := selectLanguage(lang)
	affichage(green + bold + selectLanguage["welcomeMsg"] + " \n " + reset)
	fmt.Println("")
	affichage(green + selectLanguage["rule1"] + " \n " + reset)
	fmt.Println("")
	fmt.Println("")
	affichage(green + selectLanguage["rule2"] + " \n " + reset)
	fmt.Println("")
	fmt.Println("")
	affichage(green + selectLanguage["username"] + reset)
	fmt.Scanln(&user)
	fmt.Println("")
	affichage(green + bold + selectLanguage["welcome"] + user + " ! \n" + reset)
	fmt.Println("")
	affichage(green + selectLanguage["choise"] + reset)
	fmt.Println("")
	affichage(green + selectLanguage["easy"] + " \n " + reset)
	affichage(green + selectLanguage["medium"] + " \n " + reset)
	affichage(green + selectLanguage["hard"] + " \n " + reset)
	fmt.Println("")
	affichage(green + selectLanguage["choise2"] + reset)
	fmt.Scanln(&choise)

	for choise != selectLanguage["easy"] && choise != selectLanguage["medium"] && choise != selectLanguage["hard"] {
		fmt.Println("")
		affichage(green + selectLanguage["incorrect"] + reset)
		fmt.Println("")
		affichage(green + selectLanguage["choise2"] + reset)
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
	var point int

	word := readFile(choise, lang)

	fmt.Println(string(word))
	for i := 0; i < len(word); i++ {
		underscore = append(underscore, '_')
	}
	for remainLife != 0 {
		count++
		fmt.Println(green + "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" + reset)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		affichage(green + string(underscore) + reset)
		fmt.Print("\n")
		fmt.Print("\n")
		affichage(green + selectLanguage["guessMsg"] + reset)
		fmt.Scanln(&letter)
		fmt.Println("")

		fmt.Println(remainLife)

		if letter == "menu" {
			returnMenu(letter, lang, green, reset)
		}

		if letter == selectLanguage["clue"] {
			clue(letter, lang, choise, underscore, word, green, reset, interupEasy, interupMedium, interupHard, remainLife)
		}

		for i := 0; i < len(word); i++ {
			if string(word[i]) == letter {
				underscore[i] = word[i]
				index = i
				point = point + 100
			}

		}

		if letter == "debug" || letter == "exit" || letter == "menu" || letter == selectLanguage["clue"] {
			remainLife = remainLife * 1
		} else if string(word[index]) != letter {

			if point <= 0 {
				point = 0
			} else {
				point = point - 70
			}

			life = life + 1
			lign := life*7 + life
			remainLife--
			remainLifeStr = strconv.Itoa(remainLife)
			affichage(green + selectLanguage["you have"] + " " + remainLifeStr + " " + selectLanguage["attempt"] + reset)
			fmt.Println("")

			for i := lign; i < lign+7; i++ {
				fmt.Println(green + readHangman()[i] + reset)
			}

			fmt.Println("")
		}

		if word[index] == rune(letter[0]) && remainLife != 10 {
			affichage(green + selectLanguage["you have"] + " " + remainLifeStr + " " + selectLanguage["attempt"] + reset)
			fmt.Println("")
		}

		if remainLife == 0 {
			if point <= 0 {
				point = 0
			} else {
				point = point - 200
			}
		}

		if string(word) == string(underscore) {
			point = point + 500
			affichage(green + selectLanguage["winMsg"] + " \n " + reset)
			affichage(green + selectLanguage["theWord"] + " \n " + string(word) + reset)
			break
		}

		if letter == "debug" {
			debug(letter, lang, remainLifeStr, green, reset, interupEasy, interupMedium, interupHard, choise, string(word), point, remainLife)
		}

	}

	if count == 10 {
		affichage(green + selectLanguage["finalAttempt"] + string(word) + "\n" + reset)

	}

	totalPoint := pointTot(choise, point)

	fmt.Println("")
	affichage(green + selectLanguage["score"] + "\n" + reset)
	fmt.Println("")
	affichage(green + user + " :               " + reset)
	affichage(green + strconv.Itoa(totalPoint) + " " + selectLanguage["point"] + reset)

}

func readFile(r string, lang string) []rune {

	var word []rune
	var list []string
	var file *os.File
	var err error

	selectLanguage := selectLanguage(lang)

	if lang == "fr" {
		if r == selectLanguage["easy"] {
			file, err = os.Open("facile.txt")
		} else if r == selectLanguage["medium"] {
			file, err = os.Open("moyen.txt")
		} else if r == selectLanguage["hard"] {
			file, err = os.Open("difficile.txt")
		}
	} else if lang == "en" {
		if r == selectLanguage["easy"] {
			file, err = os.Open("facileEng.txt")
		} else if r == selectLanguage["medium"] {
			file, err = os.Open("moyenEng.txt")
		} else if r == selectLanguage["hard"] {
			file, err = os.Open("difficileEng.txt")
		}
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

func pointTot(c string, p int) int {
	if c == "easy" {
		p = p * 1
	} else if c == "medium" {
		p = p * 2
	} else if c == "hard" {
		p = p * 3
	}
	return p
}

func selectLanguage(lang string) map[string]string {
	languages := make(map[string]map[string]string)

	languages["en"] = map[string]string{
		"welcomeMsg":   "Welcome to Hangman game !",
		"guessMsg":     "enter a letter:",
		"winMsg":       "You win! ",
		"loseMsg":      "Sorry, you have run out of attempts. The word was : ",
		"rule1":        "rules are simple you must find the aleatory word that the programme will choose. You have 10 attempt to find it and if you don't get it at the end you loose.",
		"rule2":        "if you want to leave the game you can tape 'menu' and you will be able to go back to the main menu. Moreover you can use clue that to move betwen 2 and 4 according to the difficulty you choose before. 2 for the most easiest and 4 for the hardest.",
		"username":     "Enter your username : ",
		"welcome":      "welcome ",
		"choise":       "you have the choise betwen 3 dificulty :",
		"easy":         "easy",
		"medium":       "medium",
		"hard":         "hard",
		"choise2":      "please make your choise : ",
		"incorrect":    "incorrect input please retry : ",
		"attempt":      "attempt left",
		"you have":     "you have",
		"clue":         "clue",
		"allClue":      "you have use all your clues !",
		"menu":         "are you sure you want to go back to the main menu ? [y/n] : ",
		"backMenu":     "ok let's go back to the main menu !",
		"continue":     "ok let's continue !",
		"lose":         "You lose! Try again",
		"finalAttempt": "you have use your 10 attempts, you lose ! the word was : ",
		"score":        "table of score :",
		"point":        "points",
		"theWord":      "The word was : ",
		"admin":        "debug mod",
		"password":     "enter the password : ",
		"exit":         " to exit the debug mod tape 'exit' : ",
		"enter":        " nothing was enter ! please try again : ",
	}

	languages["fr"] = map[string]string{
		"welcomeMsg":   "Bienvenue dans hangman!",
		"guessMsg":     "entrez une lettre : ",
		"winMsg":       "vous avez gagn\u00e9 ! ",
		"loseMsg":      "D\u00e9sol\u00e9, vous avez \u00e9puis\u00e9 vos essais. Le mot \u00e9tait :",
		"rule1":        "les r\u00e8gles sont simples, vous devez trouver le mot al\u00e9atoire que le programme choisira. Vous avez 10 essais pour le trouver et si vous ne le trouvez pas \u00e0 la fin, vous perdez.",
		"rule2":        "si vous voulez quitter le jeu, vous pouvez taper 'menu' et vous pourrez revenir au menu principal. De plus, vous pouvez utiliser un indice pour vous d\u00e9placer entre 2 et 4 en fonction de la difficult\u00e9 que vous avez choisie auparavant. 2 pour le plus facile et 4 pour le plus difficile.",
		"username":     "Entrez votre nom d'utilisateur : ",
		"welcome":      "bienvenue ",
		"choise":       "vous avez le choix entre 3 difficult\u00e9s :",
		"easy":         "facile",
		"medium":       "moyen",
		"hard":         "difficile",
		"choise2":      "faites votre choix : ",
		"incorrect":    "entr\u00e9e incorrecte veuillez r\u00e9essayer : ",
		"attempt":      "essai restant",
		"you have":     "il vous reste",
		"clue":         "indice",
		"allClue":      "vous avez utilis\u00e9 tous vos indices !",
		"menu":         "\u00eates-vous s\u00fbr de vouloir revenir au menu principal ? [y/n] : ",
		"backMenu":     "ok allons-y retour au menu principal !",
		"continue":     "ok continuons !",
		"lose":         "vous avez perdu! R\u00e9essayer",
		"finalAttempt": "vous avez utilis\u00e9 vos 10 essais, vous perdez ! le mot \u00e9tait : ",
		"score":        "tableau des scores :",
		"point":        "points",
		"theWord":      "Le mot \u00e9tait : ",
		"admin":        "mode debug",
		"password":     "entrez le mot de passe : ",
		"exit":         " pour quitter le mode debug tapez 'exit' : ",
		"enter":        " rien n'a \u00e9t\u00e9 entr\u00e9 ! veuillez r\u00e9essayer : ",
	}

	return languages[lang]
}

func debug(letter string, lang string, remainLifeStr string, green string, reset string, interupEasy int, interupMedium int, interupHard int, choise string, word string, point int, remainLife int) {

	selectLanguage := selectLanguage(lang)
	interuptEasyStr := strconv.Itoa(interupEasy)
	interuptMediumStr := strconv.Itoa(interupMedium)
	interuptHardStr := strconv.Itoa(interupHard)
	pointStr := strconv.Itoa(point)

	if letter == "debug" {
		var password string
		var exit string

		remainLife = remainLife + 1
		affichage(green + selectLanguage["password"] + "\n" + reset)
		fmt.Scanln(&password)
		if password == "root" {
			affichage(green + selectLanguage["admin"] + " \n " + reset)
			fmt.Println("")
			fmt.Println("")
			affichage(green + selectLanguage["theWord"] + "  " + string(word) + "\n" + reset)
			affichage(green + selectLanguage["score"] + "  " + pointStr + "  " + selectLanguage["point"] + "\n" + reset)
			affichage(green + remainLifeStr + " " + selectLanguage["attempt"] + "\n" + reset)
			if choise == selectLanguage["easy"] {
				affichage(green + selectLanguage["clue"] + " : " + interuptEasyStr + "\n" + reset)
			} else if choise == selectLanguage["medium"] {
				affichage(green + selectLanguage["clue"] + " : " + interuptMediumStr + "\n" + reset)
			} else if choise == selectLanguage["hard"] {
				affichage(green + selectLanguage["clue"] + " : " + interuptHardStr + "\n" + reset)
			}
			affichage(green + selectLanguage["exit"] + reset)
			fmt.Scanln(&exit)

			if exit == "exit" {
				return
			}

		}
	}
}

func returnMenu(letter string, lang string, green string, reset string) {

	selectLanguage := selectLanguage(lang)

	if letter == "menu" {
		var desire string
		affichage(green + selectLanguage["menu"] + reset)
		fmt.Scanln(&desire)
		if desire == "y" {
			affichage(green + selectLanguage["backMenu"] + " \n " + reset)
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			jeu()
		} else if desire == "n" {
			affichage(green + selectLanguage["continue"] + " \n " + reset)
		}
	}
}

func clue(letter string, lang string, choise string, underscore []rune, word []rune, green string, reset string, interupEasy int, interupMedium int, interupHard int, remainLife int) {

	selectLanguage := selectLanguage(lang)
	if letter == selectLanguage["clue"] {
		var indice int

		if choise == selectLanguage["easy"] && interupEasy <= 2 && underscore[indice] != word[indice] {
			indice = rand.Intn(len(word))
			underscore[indice] = word[indice]
			interupEasy++
			fmt.Println("boucle easy")
		} else if choise == selectLanguage["medium"] && interupMedium <= 3 && underscore[indice] != word[indice] {
			indice = rand.Intn(len(word))
			underscore[indice] = word[indice]
			interupMedium++
			fmt.Println("boucle medium")
		} else if choise == selectLanguage["hard"] && interupHard <= 4 && underscore[indice] != word[indice] {
			indice = rand.Intn(len(word))
			underscore[indice] = word[indice]
			interupHard++
			fmt.Println("boucle hard")
		} else {
			affichage(green + selectLanguage["allClue"] + "\n" + reset)
			remainLife = remainLife + 1

		}
	}

}
