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
	jsonData, err := json.Marshal(apperror)
	if err != nil {
		return err
	}
	v.writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(v.writer, string(jsonData))
	return nil
}
