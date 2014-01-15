package config

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mantishK/gonotevanilla/model"
	"log"
	"strconv"
	"strings"
)

func NewConnection() *gorp.DbMap {
	dbUserName := "root"
	dbPass := ""
	dbIp := "127.0.0.1"
	dbPortNo := 3306
	dbName := "dummy"
	if dbPass != "" {
		dbPass = ":" + dbPass
	}
	dbStringSlice := []string{dbUserName, dbPass, ":@tcp(", dbIp, ":", strconv.Itoa(dbPortNo), ")/", dbName}
	db, err := sql.Open("mysql", strings.Join(dbStringSlice, ""))
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(model.Note{}, "note").SetKeys(true, "Note_id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
