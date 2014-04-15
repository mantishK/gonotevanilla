package main

import (
	"github.com/mantishK/gonotevanilla/controllers"
	"github.com/mantishK/gonotevanilla/core/router"
	"github.com/mantishK/gonotevanilla/filters"
	"log"
	"net/http"
)

func main() {
	route()
}
func route() {
	//Create filters
	authenticateFilter := new(filters.Authenticate)
	authorizeFilter := new(filters.Authorize)

	//Create controller
	noteController := controllers.NoteController

	//Create Router
	myRouter := router.New()

	//route
	myRouter.Get("/", noteController.GetNotes, authenticateFilter, authorizeFilter)

	http.Handle("/", myRouter)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
