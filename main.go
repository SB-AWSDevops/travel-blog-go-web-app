package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func destinationsPage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/destinations.html")
}

func galleryPage(w http.ResponseWriter, r *http.Request) {
	// Render the gallery html page
	http.ServeFile(w, r, "static/gallery.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/destinations", destinationsPage)
	http.HandleFunc("/gallery", galleryPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
