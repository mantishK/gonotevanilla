package main

import (
	"github.com/mantishK/gonotevanilla/controllers"
	"github.com/mantishK/gonotevanilla/core/router"
	"log"
	"net/http"
)

func main() {
	route()
}
func route() {
	noteController := controllers.NoteController
	myRouter := router.New()
	// http.HandleFunc("/", noteController.GetNotes)
	// http.HandleFunc("/add", noteController.SaveNotes)
	myRouter.Get("/", noteController.GetNotes)
	http.Handle("/", myRouter)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
