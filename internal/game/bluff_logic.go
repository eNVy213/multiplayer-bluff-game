package game

import (
	"errors"
	"math/rand"
	"time"
)

type BluffGame struct {
	Deck        []string
	CurrentTurn int
	Players     []*Player
	Pile        []string
}

func NewBluffGame() *BluffGame {
	return &BluffGame{
		Deck:        generateDeck(),
		CurrentTurn: 0,
		Players:     make([]*Player, 0),
		Pile:        make([]string, 0),
	}
}

func generateDeck() []string {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	deck := []string{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, rank+" of "+suit)
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func (g *BluffGame) AddPlayer(player *Player) error {
	if len(g.Players) >= 4 {
		return errors.New("table is full")
	}
	g.Players = append(g.Players, player)
	return nil
}

func (g *BluffGame) PlayCard(playerID, card string, declaredRank string) error {
	currentPlayer := g.Players[g.CurrentTurn]
	if currentPlayer.ID != playerID {
		return errors.New("not your turn")
	}

	g.Pile = append(g.Pile, card)
	g.CurrentTurn = (g.CurrentTurn + 1) % len(g.Players)
	return nil
}

func (g *BluffGame) CallBluff(playerID string) ([]string, error) {
	if len(g.Pile) == 0 {
		return nil, errors.New("no cards to challenge")
	}

	lastPlayer := g.Players[(g.CurrentTurn-1+len(g.Players))%len(g.Players)]
	pile := g.Pile
	g.Pile = []string{}

	return pile, nil
}
