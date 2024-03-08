package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about.html")
}

func otherHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "other.html")
}

func main() {

	// Mappage des chemins des pages aux gestionnaires correspondants
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/other", otherHandler)

	// Servir les fichiers statiques (CSS)
	fsStatic := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	fsImage := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", fsImage))

	// Démarrage du serveur sur le port 8080
	fmt.Println("Serveur démarré sur le port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
