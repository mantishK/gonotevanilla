package views

import (
	"encoding/json"
	"fmt"
	"github.com/mantishK/gonotevanilla/core/apperror"
	"net/http"
)

type view struct {
	writer http.ResponseWriter
}

func NewView(w http.ResponseWriter) view {
	return view{w}
}

func (v *view) RenderJson(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	v.writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(v.writer, string(jsonData))
	return nil
}

func (v *view) RenderErrorJson(apperror apperror.Apperror) error {
	result := make(map[string]interface{})
	result["id"] = apperror.GetIdString()
	result["message"] = apperror.GetMessage()
	if len(apperror.GetSysMesage()) > 0 {
		result["sys_message"] = apperror.GetSysMesage()
	} else {
		result["sys_message"] = nil
	}
	if len(apperror.GetField()) > 0 {
		result["field"] = apperror.GetField()
	} else {
		result["field"] = nil
	}
	result["response"] = "error"
	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	v.writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(v.writer, string(jsonData))
	return nil
}
