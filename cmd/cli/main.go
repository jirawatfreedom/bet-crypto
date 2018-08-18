package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jirawatfreedom/bet-crypto"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}
