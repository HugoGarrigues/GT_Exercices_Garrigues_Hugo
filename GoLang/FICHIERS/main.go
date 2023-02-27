package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	menu()
}

func menu() {
	var choix int
	for {
		fmt.Println("______________________________________")
		fmt.Println("|                                     |")
		fmt.Println("|       BIENVENUE A VOUS DANS LE      |")
		fmt.Println("|           GESTIONNAIRE DE           |")
		fmt.Println("|              FICHIERS               |")
		fmt.Println("|                                     |")
		fmt.Println("|   1. Récupérer le texte d'un .txt   |")
		fmt.Println("|                                     |")
		fmt.Println("|   2. Ajouter du texte dans un .txt  |")
		fmt.Println("|                                     |")
		fmt.Println("|   3. Supprimer le contenu d'un .txt |")
		fmt.Println("|                                     |")
		fmt.Println("|   4. Remplacer le contenu du .txt   |")
		fmt.Println("|                                     |")
		fmt.Println("|   5. Quitter                        |")
		fmt.Println("|_____________________________________|")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			recuptxt()
		case 2:
			ajouttxt()
		case 3:
			supprimetxt()
		case 4:
			remplacetxt()
		case 5:
			fmt.Println("A Bientot !")
			os.Exit(0)
		default:
			fmt.Println("Votre Choix n'est pas valide, veuillez réessayer")
			menu()
		}
	}
}

func recuptxt() {
	cheminFichier := "1.txt"
	contenuFichier, err := ioutil.ReadFile(cheminFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	fmt.Println(string(contenuFichier))
	menu()
}

func ajouttxt() {
	var ajout string
	cheminFichier := "1.txt"
	contenuFichier, err := ioutil.ReadFile(cheminFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	fmt.Println(string(contenuFichier))
	fmt.Println("Que veux-tu ajouter ?")
	fmt.Scan(&ajout)
	f, err := os.OpenFile(cheminFichier, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(ajout); err != nil {
		panic(err)
	}
	fmt.Println("Le texte a bien été ajouté au fichier !")
	menu()
}

func supprimetxt() {
	cheminFichier := "1.txt"
	f, err := os.OpenFile(cheminFichier, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Le fichier texte a bien été vidé !")
	menu()
}

func remplacetxt() {
	var ajout string
	cheminFichier := "1.txt"
	f, err := os.OpenFile(cheminFichier, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Le fichier a bien été vidé !")
	fmt.Println("Que voulez-vous ajouter ?")
	fmt.Scan(&ajout)
	f, err = os.OpenFile(cheminFichier, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(ajout); err != nil {
		panic(err)
	}
	fmt.Println("Le texte a bien été ajouté !")
	menu()
}
