package controllers

import (
	// "fmt"
	"github.com/mantishK/gonotevanilla/lib/validate"
	"github.com/mantishK/gonotevanilla/model"
	"github.com/mantishK/gonotevanilla/views"
	"log"
	"net/http"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, data := Init(w, r)
	note := model.Note{}
	requiredFields := []string{"start", "limit"}
	err := validate.Required(data, requiredFields)
	if err != nil {
		result := make(map[string]interface{})
		result["error"] = err.Error()
		result["response"] = "error"
		view.RenderJson(result)
		return
	}
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
