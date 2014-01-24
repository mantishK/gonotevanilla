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
	http.HandleFunc("/add", controllers.SaveNote)
	http.HandleFunc("/edit", controllers.EditNote)
	http.HandleFunc("/delete", controllers.DeleteNote)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
