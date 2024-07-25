package main

import (
	"fmt"
	"go-with-test/http-server/pkg/poker"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker!")
	fmt.Println("Type '{Name} wins' to record a win")

	game := poker.NewGame(store, poker.BlindAlerterFunc(poker.StdOutAlert))
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
