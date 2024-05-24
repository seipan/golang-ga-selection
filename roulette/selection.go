package roulette

import (
	"errors"
	"math/rand"
	"time"

	selection "github.com/seipan/golang-ga-selection"
)

type RouletteSelector struct{}

func (rs RouletteSelector) Select(population selection.Individuals) (selection.Individuals, error) {
	if len(population) == 0 {
		return nil, errors.New("集団が空です")
	}

	// 集団の総フィットネスを計算
	totalFitness := 0
	for _, ind := range population {
		totalFitness += ind.Fitness
	}

	if totalFitness == 0 {
		return nil, errors.New("集団のフィットネス合計が0です")
	}

	rand.Seed(time.Now().UnixNano())
	selected := make(selection.Individuals, 0, len(population))

	for i := 0; i < len(population); i++ {
		randomValue := rand.Intn(totalFitness)
		cumulativeSum := 0

		for _, ind := range population {
			cumulativeSum += ind.Fitness
			if randomValue < cumulativeSum {
				selected = append(selected, ind)
				break
			}
		}
	}

	return selected, nil
}
