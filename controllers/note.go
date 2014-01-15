package controllers

import (
	"github.com/mantishK/gonotevanilla/model"
	"github.com/mantishK/gonotevanilla/views"
	"log"
	"net/http"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, _ := Init(w, r)
	note := model.Note{}
	notes, len, err := note.GetNotes(dbMap)
	if err != nil {
		log.Fatalln(err)
	}
	result := make(map[string]interface{})
	result["notes"] = notes
	result["count"] = len
	result["response"] = "ok"
	view.RenderJson(result)
}
