package matchmaking

import "sync"

type MatchQueue struct {
	Players []string
	Mutex   sync.Mutex
}

func (q *MatchQueue) AddPlayer(playerID string) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.Players = append(q.Players, playerID)
}

func (q *MatchQueue) MatchPlayers() []string {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if len(q.Players) < 4 {
		return nil
	}
	match := q.Players[:4]
	q.Players = q.Players[4:]
	return match
}
