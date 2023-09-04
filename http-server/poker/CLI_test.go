package poker_test

import (
	"strings"
	"testing"
	"yeget/Go_Application/poker"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n ")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in) // can't do this for private fields
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n ")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in) // can't do this for private fields
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}