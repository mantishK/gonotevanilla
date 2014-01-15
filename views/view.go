package views

import (
	"encoding/json"
	"fmt"
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
