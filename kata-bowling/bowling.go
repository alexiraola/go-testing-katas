package main

type BowlingGame struct {
	rolls []int
}

func (game *BowlingGame) Roll(hits int) {
	game.rolls = append(game.rolls, hits)
}

func (game BowlingGame) CalculateScore() int {
	score := 0

	for _, roll := range game.rolls {
		score += roll
	}

	return score
}
