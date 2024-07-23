package poker

type CLI struct {
	playerstore PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerstore.RecordWin("Abe")
}
