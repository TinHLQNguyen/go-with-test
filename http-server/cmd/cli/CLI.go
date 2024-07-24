package main

import (
	"bufio"
	"go-with-test/http-server/pkg/poker"
	"io"
	"strings"
)

type CLI struct {
	playerstore poker.PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store poker.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerstore: store,
		in:          bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerstore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
