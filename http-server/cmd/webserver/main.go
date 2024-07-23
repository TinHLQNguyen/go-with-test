package main

import (
	"go-with-test/http-server/pkg/poker"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
