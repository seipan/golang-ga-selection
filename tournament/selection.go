package tournament

import (
	"errors"
	"math/rand"
	"time"

	selection "github.com/seipan/golang-ga-selection"
)

type TournamentSelector struct {
	TournamentSize int
}

func (ts TournamentSelector) Select(population selection.Individuals) (selection.Individuals, error) {
	if ts.TournamentSize > len(population) {
		return nil, errors.New("トーナメントサイズが集団のサイズを超えています")
	}

	rand.Seed(time.Now().UnixNano())
	selected := make(selection.Individuals, 0, len(population))

	for i := 0; i < len(population); i++ {
		tournament := make(selection.Individuals, ts.TournamentSize)
		for j := 0; j < ts.TournamentSize; j++ {
			randomIndex := rand.Intn(len(population))
			tournament[j] = population[randomIndex]
		}

		best := tournament[0]
		for _, ind := range tournament {
			if ind.Fitness > best.Fitness {
				best = ind
			}
		}
		selected = append(selected, best)
	}

	return selected, nil
}
