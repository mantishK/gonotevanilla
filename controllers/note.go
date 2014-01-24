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
	count, err := validate.RequiredParams(params, requiredFields)
	if err != nil {
		view.RenderErrorJson(apperror.NewRequiredError(err.Error(), requiredFields[count]))
		return
	}

	start, err := strconv.Atoi(params.Get("start"))
	if err != nil {
		view.RenderErrorJson(apperror.NewInvalidInputError("", err, "start"))
		return
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		view.RenderErrorJson(apperror.NewInvalidInputError("", err, "limit"))
		return
	}

	notes, len, err := note.Get(dbMap, start, limit)
	if err != nil {
		view.RenderErrorJson(apperror.NewDBError("", err))
		return
	}
	result := make(map[string]interface{})
	result["notes"] = notes
	result["count"] = len
	result["response"] = "ok"
	view.RenderJson(result)
}

func SaveNote(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, data, _ := Init(w, r)
	note := model.Note{}
	requiredFields := []string{"title", "body"}
	count, err := validate.RequiredData(data, requiredFields)
	if err != nil {
		view.RenderErrorJson(apperror.NewRequiredError(err.Error(), requiredFields[count]))
		return
	}
	note.Content = data["body"].(string)
	note.Title = data["title"].(string)
	err = note.Save(dbMap)
	if err != nil {
		view.RenderErrorJson(apperror.NewDBError("", err))
		return
	}
	result := make(map[string]interface{})
	result["response"] = "ok"
	view.RenderJson(result)
}

func EditNote(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, data, _ := Init(w, r)
	note := model.Note{}
	requiredFields := []string{"id", "title", "body"}
	count, err := validate.RequiredData(data, requiredFields)
	if err != nil {
		view.RenderErrorJson(apperror.NewRequiredError("", requiredFields[count]))
		return
	}
	note.Note_id, err = strconv.Atoi(data["id"].(string))
	if err != nil {
		view.RenderErrorJson(apperror.NewInvalidInputError("", err, "limit"))
		return
	}
	note.Content = data["body"].(string)
	note.Title = data["title"].(string)
	updatedRows, err := note.Update(dbMap)
	if err != nil {
		view.RenderErrorJson(apperror.NewDBError("", err))
		return
	}
	result := make(map[string]interface{})
	result["response"] = "ok"
	result["updatedRows"] = updatedRows
	view.RenderJson(result)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	view := views.NewView(w)
	dbMap, _, params := Init(w, r)
	note := model.Note{}
	requiredFields := []string{"id"}
	count, err := validate.RequiredParams(params, requiredFields)
	if err != nil {
		view.RenderErrorJson(apperror.NewRequiredError(err.Error(), requiredFields[count]))
		return
	}

	note_id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		view.RenderErrorJson(apperror.NewInvalidInputError("", err, "start"))
		return
	}
	note.Note_id = note_id
	deletedRows, err := note.Delete(dbMap)
	if err != nil {
		view.RenderErrorJson(apperror.NewDBError("", err))
		return
	}
	result := make(map[string]interface{})
	result["rows affected"] = deletedRows
	result["response"] = "ok"
	view.RenderJson(result)
}
