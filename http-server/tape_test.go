package main

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	//file has 12345
	file, clean := createTempFile(t, "12345")
	defer clean()
	tape := &tape{file}
	// we have 12345 we want to clear that an write abc not abc45
	tape.Write([]byte("abc")) // reverse the order of seeking and writing
	file.Seek(0, 0)           // why do we seek after writing??

	newFileContents, _ := io.ReadAll(file)
	got := string(newFileContents)
	want := "abc"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
