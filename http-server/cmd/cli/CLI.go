package main

import (
	"bufio"
	"go-with-test/http-server/pkg/poker"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type dummyAlerter struct{}

func (d *dummyAlerter) ScheduleAlertAt(duration time.Duration, amount int) {}

type CLI struct {
	playerstore poker.PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

func NewCLI(store poker.PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerstore: store,
		in:          bufio.NewScanner(in),
		alerter:     alerter,
	}
}

func (cli *CLI) PlayPoker() {
	cli.alerter.ScheduleAlertAt(5*time.Second, 1)
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
