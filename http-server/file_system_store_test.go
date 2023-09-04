package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	//reading from the file: reading league
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins":10},
			{"Name":"Chris","Wins":33}]`)
		defer cleanDatabase()
		store, _ := NewFileSystemPlayerStore(database) // instead of &FileSystemPlayerStore{database}
		got := store.GetLeague()                       // try instead of store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
		//calling a second time
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
	// reading from the file:reading from player
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins":10},
			{"Name":"Chris","Wins":33}
		]`)
		defer cleanDatabase()
		store, _ := NewFileSystemPlayerStore(database)
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})
	// writing into file: storing wins for existing players
	//file_system_store_test.go
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name":"Cleo","Wins":10},
			{"Name":"Chris","Wins":33}
		]`)
		defer cleanDatabase()
		store, _ := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)

	})

}
func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// creating a temporary file that acts as a db, the os.File implements ReadWriteSeeker
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("couldn't create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())

	}
	return tmpfile, removeFile

}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one,%v", err)
	}
}
