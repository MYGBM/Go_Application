package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface { // we need some kind of storage system after getting the player scores
	GetPlayerScore(name string) int
	RecordWin(name string)
}
type PlayerServer struct { // making PlayerServer a struct to implement handler interface
	store PlayerStore // interface polymorphism or duck typing PlayerServer can now get access to GetPlayerScore an RecordWin
}
type StubPlayerStore struct { // fake storage
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//handling GET OR POST based on http.Method
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}
func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	//Handling GET
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player) // as you can see here PlayerServer can access the GetPlayerScore method
	if score == 0 {                         // handling 404 StatusNotFoundCase
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)

}
func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	//HandlingPOST
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
