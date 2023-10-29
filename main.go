package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
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
	var choise string             // variable qui va contenir la difficulté choisie par l'utilisateur
	var user string               // variable qui va contenir le nom de l'utilisateur
	var lang string               // variable qui va contenir la langue choisie par l'utilisateur
	var green string = "\033[32m" // couleur vert
	var reset string = "\033[0m"  // couleur par défaut
	var bold string = "\033[1m"   // gras

	encoder := charmap.Windows1252.NewEncoder() // variable qui va permettre d'afficher les accents correctement

	affichage(green + "select your language [fr] [en] : " + reset) // demande a l'utilisateur de choisir sa langue
	fmt.Scanln(&lang)

	// boucle qui va permettre de faire tourner la demande en boucle  tant que l'utilisateur n'a pas choisi une langue correcte

	for lang != "fr" && lang != "en" {
		affichage(green + "incorrect answer please select your language [fr] [en] : " + reset) // demande a l'utilisateur de choisir sa langue
		fmt.Scanln(&lang)                                                                      // récupération de la langue choisie par l'utilisateur
	}

	// phase de démarage du jeu avec les règles et le choix de la difficulté
	fmt.Println("")
	selectLanguage := selectLanguage(lang) // on récupère la langue choisie par l'utilisateur
	affichage(green + bold + selectLanguage["welcomeMsg"] + " \n " + reset)
	fmt.Println("")

	// sert a afficher les accents correctement

	rule1 := selectLanguage["rule1"]                       // récupération de la règle 1
	convertRule1, _, _ := transform.String(encoder, rule1) // conversion de la règle 1 en string
	affichage(green + convertRule1 + "\n" + reset)         // affichage de la règle 1

	fmt.Println("")
	fmt.Println("")
	rule2 := selectLanguage["rule2"]
	convertRule2, _, _ := transform.String(encoder, rule2)
	affichage(green + convertRule2 + "\n" + reset)
	fmt.Println("")
	fmt.Println("")
	username := selectLanguage["username"]
	convertUsername, _, _ := transform.String(encoder, username)
	affichage(green + convertUsername + reset)
	fmt.Scanln(&user)
	fmt.Println("")
	affichage(green + bold + selectLanguage["welcome"] + user + " ! \n" + reset)
	fmt.Println("")
	choix := selectLanguage["choise"]
	convertChoix, _, _ := transform.String(encoder, choix)
	affichage(green + convertChoix + " \n " + reset)
	fmt.Println("")
	affichage(green + selectLanguage["easy"] + "\n" + reset)
	affichage(green + selectLanguage["medium"] + "\n" + reset)
	affichage(green + selectLanguage["hard"] + "\n" + reset)
	fmt.Println("")
	affichage(green + selectLanguage["choise2"] + reset)
	fmt.Scanln(&choise)

	// boucle qui va permettre de faire tourner le jeu tant que l'utilisateur n'a pas choisi une difficulté correcte

	for choise != selectLanguage["easy"] && choise != selectLanguage["medium"] && choise != selectLanguage["hard"] { // si l'utilisateur n'a pas choisi une difficulté correcte
		fmt.Println("")
		incorrect := selectLanguage["incorrect"] // on affiche que la difficulté n'est pas correcte
		convertIncorrect, _, _ := transform.String(encoder, incorrect)
		affichage(green + convertIncorrect + " " + reset) // on demande a l'utilisateur de choisir une difficulté
		fmt.Scanln(&choise)                               // on récupère la difficulté choisie par l'utilisateur
		fmt.Println("")
	}

	var remainLifeStr string  // variable qui va contenir le nombre de vie restante converti en string
	var remainLife int = 10   // variable qui va contenir le nombre de vie restante
	var life int = -1         // variable qui va contenir le nombre de vie perdu
	var letter string         // variable qui va contenir la lettre entrée par l'utilisateur
	var underscore []rune     // tableau qui va contenir les underscore qui vont s'afficher a l'écran et qui seront remplacer par les lettres trouvées
	var index int             // variable qui va contenir l'index de la lettre trouvée
	var interupEasy int = 0   // variable qui va contenir le nombre d'indice utilisé pour le mode facile
	var interupMedium int = 0 // variable qui va contenir le nombre d'indice utilisé pour le mode moyen
	var interupHard int = 0   // variable qui va contenir le nombre d'indice utilisé pour le mode difficile
	var point int             // variable qui va contenir le score

	// récupération du fichier avec les mots en fonction de la difficulté choisie

	word := readFile(choise, lang) // on récupère le mot

	// initialisation du tableau de caractère qui va contenir les underscore qui vont s'afficher a l'écran et qui seront remplacer par les lettres trouvées

	for i := 0; i < len(word); i++ { // boucle qui va parcourir la longueur du mot
		underscore = append(underscore, '_') // on ajoute un underscore dans le tableau
	}

	// boucle qui va permettre de faire tourner le jeu tant que le joueur n'a pas trouvé le mot ou qu'il n'a pas utilisé ses 10 essais

	for remainLife != 0 { // tant que l'utilisateur n'a pas utilisé ses 10 essais

		// demande a l'utilisateur d'entrer une lettre
		fmt.Println(green + "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" + reset)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		underscoreBis := string(underscore)
		convertUnderscoreBis, _, _ := transform.String(encoder, underscoreBis) // on converti le tableau de rune en string pour pouvoir l'afficher
		affichage(green + convertUnderscoreBis + reset)
		fmt.Print("\n")
		fmt.Print("\n")
		affichage(green + selectLanguage["guessMsg"] + reset) // on demande a l'utilisateur d'entrer une lettre
		fmt.Scanln(&letter)                                   // on récupère la lettre entrée par l'utilisateur
		fmt.Println("")

		/* fmt.Println(remainLife) */

		// cette fonction sert a revenir au début du jeu si l'utilisateur tape menu

		if letter == "menu" { // si l'utilisateur tape menu
			returnMenu(letter, lang, green, reset) // on appelle la fonction qui va permettre de revenir au début du jeu
		}

		// cette fonction sert a donner un indice a l'utilisateur si il tape clue

		if letter == selectLanguage["clue"] { // si l'utilisateur tape clue ou indice selon la langue
			var indice int // variable qui va contenir le nombre d'indice utilisé

			if choise == selectLanguage["easy"] && interupEasy < 2 && underscore[indice] != word[indice] { // ici on vérifie que l'utilisateur n'a pas déjà utilisé ses indices et que l'indice n'est pas déjà utilisé pour le mode facile
				indice = rand.Intn(len(word))     // ici on choisi un indice aléatoire dans le mot
				underscore[indice] = word[indice] // ici on remplace un underscore par une lettre du mot
				interupEasy++                     // ici on incrémente le nombre d'indice utilisé

			} else if choise == selectLanguage["medium"] && interupMedium < 3 && underscore[indice] != word[indice] { // ici on vérifie que l'utilisateur n'a pas déjà utilisé ses indices et que l'indice n'est pas déjà utilisé pour le mode moyen
				indice = rand.Intn(len(word))     // ici on choisi un indice aléatoire dans le mot
				underscore[indice] = word[indice] // ici on remplace un underscore par une lettre du mot
				interupMedium++                   // ici on incrémente le nombre d'indice utilisé

			} else if choise == selectLanguage["hard"] && interupHard < 4 { // ici on vérifie que l'utilisateur n'a pas déjà utilisé ses indices et que l'indice n'est pas déjà utilisé pour le mode difficile
				indice = rand.Intn(len(word))     // ici on choisi un indice aléatoire dans le mot
				underscore[indice] = word[indice] // ici on remplace un underscore par une lettre du mot
				interupHard++                     // ici on incrémente le nombre d'indice utilisé

			} else {
				// si l'utilisateur a utilisé tous ses indices il ne pourra plus en avoir
				allClue := selectLanguage["allClue"]
				convertAllClue, _, _ := transform.String(charmap.Windows1252.NewEncoder(), allClue)
				affichage(green + convertAllClue + "\n" + reset)
			}
		}

		// cette fonction sert a quitter le jeu si l'utilisateur rentre un esapce vide

		if letter == "" {
			for letter == "" { // boucle qui va permettre de faire tourner le jeu tant que l'utilisateur n'a pas entré une lettre
				entrer := selectLanguage["enter"]
				convertEntrer, _, _ := transform.String(encoder, entrer)
				affichage(green + convertEntrer + reset)
				fmt.Scanln(&letter)
				fmt.Println("")
			}
		}

		// cette fonction sert a voir si la lettre rentrée par l'utilisateur est dans le mot

		for i := 0; i < len(word); i++ { // boucle qui va parcourir le mot
			if string(word[i]) == letter { // si la lettre est dans le mot
				underscore[i] = word[i] // on remplace l'underscore par la lettre
				index = i               // on récupère l'index de la lettre
				point = point + 100     // on ajoute 100 points au score
			}

		}

		if letter == "debug" || letter == "exit" || letter == "menu" || letter == selectLanguage["clue"] || letter == "" { // si l'utilisateur rentre une commande spéciale on ne fait rien
			remainLife = remainLife * 1
		} else if string(word[index]) != letter { // si la lettre n'est pas dans le mot

			if point <= 0 { // si le score est inférieur ou égal a 0 on ne fait rien
				point = 0
			} else { // sinon on enlève 70 points au score
				point = point - 70
			}

			life = life + 1                                                                                               // on incrémente le nombre de vie perdu
			lign := life*7 + life                                                                                         // on calcule la ligne a partir de laquelle on va afficher le pendu
			remainLife--                                                                                                  // on enlève une vie
			remainLifeStr = strconv.Itoa(remainLife)                                                                      // on converti le nombre de vie restante en string
			affichage(green + selectLanguage["you have"] + " " + remainLifeStr + " " + selectLanguage["attempt"] + reset) // on affiche le nombre de vie restante
			fmt.Println("")

			for i := lign; i < lign+7; i++ { // boucle qui va afficher le pendu
				fmt.Println(green + readHangman()[i] + reset) // on affiche le pendu
			}

			fmt.Println("")
		}

		if word[index] == rune(letter[0]) && remainLife != 10 { // si la lettre est dans le mot et que l'utilisateur n'a pas utilisé ses 10 essais
			affichage(green + selectLanguage["you have"] + " " + remainLifeStr + " " + selectLanguage["attempt"] + reset)
			fmt.Println("")
		}

		if remainLife == 0 { // si l'utilisateur a utilisé ses 10 essais
			if point <= 0 {
				point = 0
			} else {
				point = point - 200
			}
		}

		// cette fonction sert a voir si l'utilisateur a gagné

		if string(word) == string(underscore) { // si le mot est trouvé
			point = point + 500 // on ajoute 500 points au score
			winMsg := selectLanguage["winMsg"]
			convertWinMsg, _, _ := transform.String(encoder, winMsg)
			affichage(green + convertWinMsg + "\n" + reset) // on affiche que l'utilisateur a gagné
			theWord := selectLanguage["theWord"]
			wordBis := word
			convertWordBis, _, _ := transform.String(encoder, string(wordBis))
			convertTheWord, _, _ := transform.String(encoder, theWord)
			affichage(green + convertTheWord + convertWordBis + "\n" + reset) // on affiche le mot
			break
		}

		// cette fonction sert a afficher les propriétés du jeu pour pouvoir débuger plus facilement

		if letter == "debug" { // si l'utilisateur rentre debug
			debug(letter, lang, remainLifeStr, green, reset, interupEasy, interupMedium, interupHard, choise, string(word), point, remainLife) // on appelle la fonction qui va afficher les propriétés du jeu
		}

	}

	// cette fonction s'affiche si l'utilisateur a perdu car il n'a plus d'essais

	if remainLife == 0 { // si l'utilisateur a utilisé ses 10 essais
		loseMsg := selectLanguage["loseMsg"]
		convertLoseMsg, _, _ := transform.String(encoder, loseMsg)
		affichage(green + convertLoseMsg + " " + string(word) + "\n" + reset) // on affiche que l'utilisateur a perdu
	}

	totalPoint := pointTot(choise, point) // on calcule le score total

	fmt.Println("")
	affichage(green + selectLanguage["score"] + "\n" + reset) // on affiche le score
	fmt.Println("")
	affichage(green + user + " :               " + reset)                               // on affiche le nom de l'utilisateur
	affichage(green + strconv.Itoa(totalPoint) + " " + selectLanguage["point"] + reset) // on affiche le score total

}

func readFile(r string, lang string) []rune {

	var word []rune   // tableau qui va contenir le mot
	var list []string // tableau qui va contenir les mots du fichier
	var file *os.File // fichier qui va contenir les mots
	var err error     // variable qui va contenir les erreurs

	selectLanguage := selectLanguage(lang) // on récupère la langue choisie par l'utilisateur

	if lang == "fr" { // on récupère le fichier en fonction de la langue choisie et de la difficulté choisie
		if r == selectLanguage["easy"] {
			file, err = os.Open("facile.txt") // on ouvre le fichier
		} else if r == selectLanguage["medium"] {
			file, err = os.Open("moyen.txt") // on ouvre le fichier
		} else if r == selectLanguage["hard"] {
			file, err = os.Open("difficile.txt") // on ouvre le fichier
		}
	} else if lang == "en" {
		if r == selectLanguage["easy"] {
			file, err = os.Open("facileEng.txt") // on ouvre le fichier
		} else if r == selectLanguage["medium"] {
			file, err = os.Open("moyenEng.txt") // on ouvre le fichier
		} else if r == selectLanguage["hard"] {
			file, err = os.Open("difficileEng.txt") // on ouvre le fichier
		}
	}

	// si il y a une erreur on l'affiche

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // on ferme le fichier

	// on lit le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	// on vérifie qu'il n'y a pas d'erreur
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	randIndex := rand.Intn(len(list)) // on choisi un mot aléatoire dans le fichier
	word = []rune(list[randIndex])    // on converti le mot en tableau de rune
	return word                       // on retourne le mot

}

func readHangman() []string { // fonction qui va lire le fichier contenant le pendu
	var hangman []string // tableau qui va contenir le pendu

	file, err := os.Open("hangman.txt") // on ouvre le fichier

	// on verifie qu'il n'y a pas d'erreur
	if err != nil {
		log.Fatal(err)
	}

	// on ferme le fichier
	defer file.Close()

	// on lit le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hangman = append(hangman, scanner.Text())
	}

	// on verifie qu'il n'y a pas d'erreur

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return hangman // on retourne le pendu
}

func draw(min int, max int) []string { // fonction qui va permettre de dessiner le pendu
	tab := readHangman() // on récupère le pendu
	var newtab []string  // on initialise un nouveau tableau qui va contenir le pendu

	for i := min; i <= max; i++ { // boucle qui va parcourir le pendu
		newtab = append(newtab, tab[i]) // on ajoute les lignes du pendu dans le nouveau tableau
	}
	return newtab // on retourne le nouveau tableau
}

func affichage(n string) { // fonction qui va permettre d'afficher le pendu
	for i := 0; i < len(n); i++ { // boucle qui va parcourir le pendu
		fmt.Print(string(n[i]))                         // on affiche le pendu
		time.Sleep(time.Duration(1) * time.Millisecond) // on met un délai de 1 milliseconde entre chaque ligne du pendu
	}
}

func pointTot(c string, p int) int { // fonction qui va calculer le score total
	if c == "easy" {
		p = p * 1
	} else if c == "medium" {
		p = p * 2
	} else if c == "hard" {
		p = p * 3
	}
	return p
}

func selectLanguage(lang string) map[string]string { // fonction qui va permettre de choisir la langue
	languages := make(map[string]map[string]string) // on initialise un tableau qui va contenir les langues

	languages["en"] = map[string]string{ // on ajoute les langues dans le tableau
		"welcomeMsg":   "Welcome to Hangman game !",
		"guessMsg":     "enter a letter : ",
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
		"exit":         "to exit the debug mod tape 'exit' : ",
		"enter":        "nothing was enter ! please try again : ",
	}

	languages["fr"] = map[string]string{ // on ajoute les langues dans le tableau
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
		"exit":         "pour quitter le mode debug tapez 'exit' : ",
		"enter":        "rien n'a \u00e9t\u00e9 entr\u00e9 ! veuillez r\u00e9essayer : ",
	}

	return languages[lang] // on retourne la langue choisie
}

func debug(letter string, lang string, remainLifeStr string, green string, reset string, interupEasy int, interupMedium int, interupHard int, choise string, word string, point int, remainLife int) { // fonction qui va afficher les propriétés du jeu pour pouvoir débuger plus facilement

	selectLanguage := selectLanguage(lang)           // on récupère la langue choisie par l'utilisateur
	interuptEasyStr := strconv.Itoa(interupEasy)     // on converti le nombre d'indice utilisé en string pour pouvoir l'afficher plus tard
	interuptMediumStr := strconv.Itoa(interupMedium) // on converti le nombre d'indice utilisé en string pour pouvoir l'afficher plus tard
	interuptHardStr := strconv.Itoa(interupHard)     // on converti le nombre d'indice utilisé en string pour pouvoir l'afficher plus tard
	pointStr := strconv.Itoa(point)                  // on converti le score en string pour pouvoir l'afficher plus tard

	if letter == "debug" { // si l'utilisateur rentre debug
		var password string
		var exit string

		remainLife = remainLife + 1                                  // on ajoute une vie a l'utilisateur pour qu'il puisse voir le pendu
		affichage(green + selectLanguage["password"] + "\n" + reset) // on demande a l'utilisateur d'entrer le mot de passe
		fmt.Scanln(&password)                                        // on récupère le mot de passe
		if password == "root" {                                      // si le mot de passe est le bon on affiche les propriétés du jeu
			affichage(green + selectLanguage["admin"] + " \n " + reset)
			fmt.Println("")
			fmt.Println("")
			theWord := selectLanguage["theWord"]
			wordBis := word
			convertWordBis, _, _ := transform.String(charmap.Windows1252.NewEncoder(), string(wordBis))
			convertTheWord, _, _ := transform.String(charmap.Windows1252.NewEncoder(), theWord)
			affichage(green + convertTheWord + "  " + convertWordBis + "\n" + reset)
			affichage(green + selectLanguage["score"] + "  " + pointStr + "  " + selectLanguage["point"] + "\n" + reset)
			affichage(green + remainLifeStr + " " + selectLanguage["attempt"] + "\n" + reset)
			if choise == selectLanguage["easy"] { // si la difficulté est facile on affiche le nombre d'indice utilisé
				affichage(green + selectLanguage["clue"] + " : " + interuptEasyStr + "\n" + reset)
			} else if choise == selectLanguage["medium"] { // si la difficulté est moyenne on affiche le nombre d'indice utilisé
				affichage(green + selectLanguage["clue"] + " : " + interuptMediumStr + "\n" + reset)
			} else if choise == selectLanguage["hard"] { // si la difficulté est difficile on affiche le nombre d'indice utilisé
				affichage(green + selectLanguage["clue"] + " : " + interuptHardStr + "\n" + reset)
			}
			affichage(green + selectLanguage["exit"] + reset) // on demande a l'utilisateur de taper exit pour quitter le mode debug
			fmt.Scanln(&exit)                                 // on récupère la commande de l'utilisateur

			if exit == "exit" { // si l'utilisateur tape exit on quitte le mode debug
				return
			}

		}
	}
}

func returnMenu(letter string, lang string, green string, reset string) { // fonction qui va permettre de revenir au début du jeu si l'utilisateur tape menu

	selectLanguage := selectLanguage(lang)

	if letter == "menu" { // si l'utilisateur tape menu
		var desire string
		menu := selectLanguage["menu"]
		convertMenu, _, _ := transform.String(charmap.Windows1252.NewEncoder(), menu)
		affichage(green + convertMenu + reset) // on demande a l'utilisateur si il veut revenir au menu
		fmt.Scanln(&desire)                    // on récupère la réponse de l'utilisateur
		if desire == "y" {
			backMenu := selectLanguage["backMenu"]
			convertBackMenu, _, _ := transform.String(charmap.Windows1252.NewEncoder(), backMenu)
			affichage(green + convertBackMenu + " \n " + reset) // on affiche que l'utilisateur va revenir au menu
			cmd := exec.Command("cmd", "/c", "cls")             // on efface l'écran
			cmd.Stdout = os.Stdout                              // on affiche l'écran
			cmd.Run()                                           // on lance la commande
			jeu()                                               // on relance le jeu
		} else if desire == "n" {
			affichage(green + selectLanguage["continue"] + " \n " + reset)
		}
	}
}
