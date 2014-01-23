package controllers

import (
	// "fmt"
	"github.com/mantishK/gonotevanilla/core/apperror"
	"github.com/mantishK/gonotevanilla/core/validate"
	"github.com/mantishK/gonotevanilla/model"
	"github.com/mantishK/gonotevanilla/views"
	"net/http"
	"strconv"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, _, params := Init(w, r)
	note := model.Note{}
	requiredFields := []string{"start", "limit"}
	// err := validate.RequiredData(data, requiredFields)
	count, err := validate.RequiredParams(params, requiredFields)
	if err != nil {
		view.RenderErrorJson(apperror.NewRequiredError(err.Error(), requiredFields[count]))
		// result := make(map[string]interface{})
		// result["error"] = err.Error()
		// result["response"] = "error"
		// view.RenderJson(result)
		return
	}

	start, err := strconv.Atoi(params.Get("start"))
	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		result := make(map[string]interface{})
		result["error"] = err.Error()
		result["response"] = "error"
		view.RenderJson(result)
		return
	}

	notes, len, err := note.Get(dbMap, start, limit)
	if err != nil {
		result := make(map[string]interface{})
		result["error"] = err.Error()
		result["response"] = "error"
		view.RenderJson(result)
		return
	}
	result := make(map[string]interface{})
	result["notes"] = notes
	result["count"] = len
	result["response"] = "ok"
	view.RenderJson(result)
}

func SaveNotes(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, data, _ := Init(w, r)
	note := model.Note{}
	requiredFields := []string{"title", "body"}
	count, err := validate.RequiredData(data, requiredFields)
	//err := validate.RequiredParams(params, requiredFields)
	if err != nil {
		view.RenderErrorJson(apperror.NewRequiredError(err.Error(), requiredFields[count]))
		// result := make(map[string]interface{})
		// result["error"] = err.Error()
		// result["response"] = "error"
		// view.RenderJson(result)
		return
	}
	note.Content = data["body"].(string)
	note.Title = data["title"].(string)
	err = note.Save(dbMap)
	if err != nil {
		result := make(map[string]interface{})
		result["error"] = err.Error()
		result["response"] = "error"
		view.RenderJson(result)
		return
	}
	result := make(map[string]interface{})
	result["response"] = "ok"
	view.RenderJson(result)
}
