package main

import (
	"github.com/mantishK/gonotevanilla/controllers"
	"log"
	"net/http"
)

func main() {
	route()
}
func route() {
	http.HandleFunc("/", controllers.GetNotes)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
