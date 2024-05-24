package ranking

import (
	"errors"
	"math/rand"
	"sort"
	"time"

	selection "github.com/seipan/golang-ga-selection"
)

type RankingSelector struct{}

func (rs RankingSelector) Select(population selection.Individuals) (selection.Individuals, error) {
	if len(population) == 0 {
		return nil, errors.New("集団が空です")
	}

	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})

	totalRank := 0
	rankWeights := make([]int, len(population))
	for i := range population {
		rankWeights[i] = len(population) - i
		totalRank += rankWeights[i]
	}

	rand.Seed(time.Now().UnixNano())
	selected := make(selection.Individuals, 0, len(population))

	for i := 0; i < len(population); i++ {
		randomValue := rand.Intn(totalRank)
		cumulativeSum := 0

		for j, weight := range rankWeights {
			cumulativeSum += weight
			if randomValue < cumulativeSum {
				selected = append(selected, population[j])
				break
			}
		}
	}

	return selected, nil
}
