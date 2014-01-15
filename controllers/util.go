package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/mantishK/gonotevanilla/config"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) (*gorp.DbMap, interface{}) {
	dbMap := config.NewConnection()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var data interface{}
	json.Unmarshal(buf.Bytes(), &data)
	fmt.Println(data)
	return dbMap, data
}
