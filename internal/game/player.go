package game

import "github.com/eNVy213/multiplayer-bluff-game/internal/network"

type Card struct {
	Value string
	Suit  string
}

type Player struct {
	ID         string `json:"id"`
	Hand       []Card `json:"hand"`
	IsBluffing bool
	Score      int
	Name       string              `json:"name"`
	Connection *network.Connection `json:"-"`
}

func NewPlayer(id string, hand []Card) *Player {
	return &Player{
		ID:    id,
		Hand:  hand,
		Score: 0,
	}
}

func (p *Player) PlayCard(card Card) bool {
	for i, c := range p.Hand {
		if c == card {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			return true
		}
	}
	return false
}

func (p *Player) AddCards(cards []Card) {
	p.Hand = append(p.Hand, cards...)
}

func (p *Player) HasCards() bool {
	return len(p.Hand) > 0
}
