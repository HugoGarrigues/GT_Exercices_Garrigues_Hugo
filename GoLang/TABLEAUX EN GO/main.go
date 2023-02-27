package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Lien vers le fichier CSS //
	static := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", static))
	// Changement de seed pour avoir des nombres aléatoires diffèrents a chaque démarrage du serveur//
	rand.Seed(time.Now().UnixNano())
	// Création du tableau a 2 dimensions avec des nombres aléatoires//
	aleatoire1 := rand.Intn(10) + 3
	aleatoire2 := rand.Intn(10) + 3
	tableau := make([][]int, aleatoire1)
	for i := range tableau {
		tableau[i] = make([]int, aleatoire2)
		for a := range tableau[i] {
			tableau[i][a] = rand.Intn(100)
		}
	}
	// Lien vers le fichier html //
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	// Création du serveur //
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, tableau)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
