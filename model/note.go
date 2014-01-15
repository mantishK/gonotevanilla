package model

import (
	"github.com/coopernurse/gorp"
)

type Note struct {
	Note_id  int    `db:"note_id"`
	Content  string `db:"note"`
	Title    string `db:"title"`
	Created  int64  `db:"created"`
	Modified int64  `db:"modified"`
}

func (n *Note) GetNotes(dbMap *gorp.DbMap) ([]Note, int, error) {
	notes := []Note{}
	_, err := dbMap.Select(&notes, "SELECT * FROM note ORDER BY modified DESC")
	if err != nil {
		return nil, 0, err
	}
	return notes, len(notes), nil
}
