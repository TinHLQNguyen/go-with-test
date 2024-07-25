package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt         = "Please enter the number of player: "
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
	BadWinnerInputErrMsg = "Bad value received for winner, please try again with a number"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayer, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayer)

	userInput := cli.readLine()
	winner, err := extractWinner(userInput)
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputErrMsg)
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	winner, ok := strings.CutSuffix(userInput, " wins")
	if !ok {
		return "", errors.New("")
	}
	return winner, nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
