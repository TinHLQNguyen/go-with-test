package poker

import "time"

type Game interface {
	Start(numberOfPlayer int)
	Finish(winner string)
}

type TexasHoldem struct {
	playerstore PlayerStore
	alerter     BlindAlerter
}

func NewGame(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{
		playerstore: store,
		alerter:     alerter,
	}
}

func (g *TexasHoldem) Start(numberOfPlayer int) {
	blindIncrement := time.Duration(5+numberOfPlayer) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindIncrement
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.playerstore.RecordWin(winner)
}
