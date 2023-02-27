package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	static := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", static))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			fmt.Println("Nom:", r.Form.Get("nom"))
			fmt.Println("Prénom:", r.Form.Get("prenom"))
		}
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
