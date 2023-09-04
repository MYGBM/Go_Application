package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}

	}
	return nil
}

// parses the database and creates a new object league
// decodes the json into normal ds, stores it in "database "
func NewLeague(rdr io.Reader) (league League, err error) {
	err = json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)

	}
	return
}
