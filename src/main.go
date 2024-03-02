package main

import (
	"log"
	"net/http"

	pages "github.com/kjblanchard/sgCom/pages"
)



func main() {
	pages.LoadTemplates()
	http.HandleFunc("/", pages.HomepageHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
