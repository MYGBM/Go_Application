package main

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
} // This is the actual storage mechanism unlike StubPlayerStore

//initialize the InMemoryPlayerStore struct
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}

}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
