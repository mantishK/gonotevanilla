package model

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Note struct {
	Note_id  int    `db:"note_id"`
	Content  string `db:"note"`
	Title    string `db:"title"`
	Created  int64  `db:"created"`
	Modified int64  `db:"modified"`
}

func (n *Note) Get(dbMap *gorp.DbMap, start, limit int) ([]Note, int, error) {
	notes := []Note{}
	_, err := dbMap.Select(&notes, "SELECT * FROM note ORDER BY modified DESC LIMIT ?,?", start, limit)
	if err != nil {
		return nil, 0, err
	}
	return notes, len(notes), nil
}

func (n *Note) Save(dbMap *gorp.DbMap) error {
	n.Created = time.Now().Unix()
	n.Modified = time.Now().Unix()
	err := dbMap.Insert(n)
	if err != nil {
		return err
	}
	return nil
}
