package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// kept empty and hard-coded response until one is implemented
type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := i.store[name]
	return score, ok
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
