package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/coopernurse/gorp"
	"github.com/mantishK/gonotevanilla/config"
	"net/http"
	"net/url"
)

func Init(w http.ResponseWriter, r *http.Request) (*gorp.DbMap, map[string]interface{}, url.Values) {
	dbMap := config.NewConnection()
	var data interface{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	json.Unmarshal(buf.Bytes(), &data)
	if data == nil {
		return dbMap, nil, r.URL.Query()
	}
	return dbMap, data.(map[string]interface{}), r.URL.Query()
}
