package main

import (
	"fmt"
)

func main() {
	menu()
}

func menu() {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|    BIENVENUE A VOUS DANS LE JEU     |")
	fmt.Println("|           DES ALLUMETTES            |")
	fmt.Println("|                                     |")
	fmt.Println("|   Choisissez le nombre d'allumettes |")
	fmt.Println("|    (le nombre d'allumettes doit     |")
	fmt.Println("|         être supèrieur a 4 )        |")
	fmt.Println("|_____________________________________|")
	fmt.Scan(&choix)
	if choix < 4 {
		fmt.Println("Le nombre d'allumettes doit être supèrieur a 4")
		menu()
	} else {
		fmt.Println("Il y aura donc ", choix, "allumettes")
		fmt.Println("")
		jeu(choix)
	}
}

func jeu(choix int) {
	var Joueur1 int
	var Joueur2 int
	for choix > 0 {
		fmt.Println("C'est à ton tour joueur 1 choisis le nombre d'allumettes que tu souhaites (entre 1 et 3)")
		fmt.Scan(&Joueur1)
		if Joueur1 < 1 || Joueur1 > 3 {
			fmt.Println("Non, tu peux uniquement choisir entre 1 et 3 allumettes")
			jeu(choix)
		} else {
			choix = choix - Joueur1
			fmt.Println("Il reste", choix, "allumettes")
			if choix == 0 {
				fmt.Println("Joueur 1 tu as perdu !")
			} else {
				fmt.Println("C'est à ton tour joueur 2 choisis le nombre d'allumettes que tu souhaites (entre 1 et 3)")
				fmt.Scan(&Joueur2)
				if Joueur2 > 3 || Joueur2 < 1 {
					fmt.Println("Non, tu peux uniquement choisir entre 1 et 3 allumettes")
					jeu(choix)
				} else {
					choix = choix - Joueur2
					fmt.Println("Il reste", choix, "allumettes")
					if choix == 0 {
						fmt.Println("Joueur 2 tu as perdu")
					}
				}
			}
		}
	}
}
