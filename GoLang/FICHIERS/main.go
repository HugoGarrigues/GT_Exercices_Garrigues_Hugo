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
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Votre Choix n'est pas valide, veuillez réessayer")
			menu()
		}
	}
}

func recuptxt() {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|      Quel fichier voulez-vous ?     |")
	fmt.Println("|                                     |")
	fmt.Println("|           1. Fichier 1              |")
	fmt.Println("|           2. Fichier 2              |")
	fmt.Println("|           3. Fichier 3              |")
	fmt.Println("|_____________________________________|")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		cheminFichier := "1.txt"
		contenuFichier, err := ioutil.ReadFile(cheminFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println(string(contenuFichier))
		menu()
	case 2:
		cheminFichier := "2.txt"
		contenuFichier, err := ioutil.ReadFile(cheminFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println(string(contenuFichier))
		menu()
	case 3:
		cheminFichier := "3.txt"
		contenuFichier, err := ioutil.ReadFile(cheminFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println(string(contenuFichier))
		menu()
	default:
		fmt.Println("Votre Choix n'est pas valide, veuillez réessayer")
		recuptxt()
	}
}

func ajouttxt() {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|      Quel fichier voulez-vous ?     |")
	fmt.Println("|                                     |")
	fmt.Println("|           1. Fichier 1              |")
	fmt.Println("|           2. Fichier 2              |")
	fmt.Println("|           3. Fichier 3              |")
	fmt.Println("|_____________________________________|")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		cheminFichier := "1.txt"
		contenuFichier, err := ioutil.ReadFile(cheminFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println(string(contenuFichier))
		var ajout string
		fmt.Println("Que voulez-vous ajouter ?")
		fmt.Scan(&ajout)
		f, err := os.OpenFile(cheminFichier, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(ajout); err != nil {
			panic(err)
		}
		fmt.Println("Le texte a bien été ajouté !")
		menu()
	case 2:
		cheminFichier := "2.txt"
		contenuFichier, err := ioutil.ReadFile(cheminFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println(string(contenuFichier))
		var ajout string
		fmt.Println("Que voulez-vous ajouter ?")
		fmt.Scan(&ajout)
		f, err := os.OpenFile(cheminFichier, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(ajout); err != nil {
			panic(err)
		}
		fmt.Println("Le texte a bien été ajouté !")
		menu()
	case 3:
		cheminFichier := "3.txt"
		contenuFichier, err := ioutil.ReadFile(cheminFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}
		fmt.Println(string(contenuFichier))
		var ajout string
		fmt.Println("Que voulez-vous ajouter ?")
		fmt.Scan(&ajout)
		f, err := os.OpenFile(cheminFichier, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(ajout); err != nil {
			panic(err)
		}
		fmt.Println("Le texte a bien été ajouté !")
		menu()
	default:
		fmt.Println("Votre Choix n'est pas valide, veuillez réessayer")
		ajouttxt()
	}
}

func supprimetxt() {
}

func remplacetxt() {
}
