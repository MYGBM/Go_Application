package main

import (
	"fmt"
	"log"
	"os"
	"yeget/Go_Application/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win ")
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()

}
