package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder // using io.Reader when called again will have nothing to read
	league   League        //go to league.go, it initialises a player using league
}

// NewFileSystemPlayerStore is a constructor function for creating a new instance of the FileSystemPlayerStore
// returning *FileSystemPlayerStore allows to create a new instance of the FileSystemPlayerStore
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := intitialisePlayerDBFile(file)
	if err != nil { // the file may be empty
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}
	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s,%v", file.Name(), err)
	}
	return &FileSystemPlayerStore{ // this is where a new instance of the FileSystemPlayerStore is created
		database: json.NewEncoder(&tape{file}), // writes json data to the file, where  do we get it ? check at home and tape.go
		league:   league,
	}, nil
}
func intitialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0) // seek to the beginning
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file from file %s,%v,", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)

	}
	return nil
}

// The FileSystem PlayerStore implements the PlayerStore interface by calling the three methods of PlayerStore( GetLeague(), GetPlayerScore(), RecordWin())
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league

}
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	// player := f.GetLeague().Find(name) test this way
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name) // instead of f.GetLeague().Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)

}
