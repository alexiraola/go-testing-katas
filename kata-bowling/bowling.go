package main

const MaxScore = 10

type ScoreCalculator interface {
	shift() int
	calculate(game BowlingGame, index int) int
}

type StrikeCalculator struct{}

func (c StrikeCalculator) shift() int {
	return 1
}

func (c StrikeCalculator) calculate(game BowlingGame, index int) int {
	return MaxScore + game.rolls[index+1] + game.rolls[index+2]
}

type SpareCalculator struct{}

func (c SpareCalculator) shift() int {
	return 2
}

func (c SpareCalculator) calculate(game BowlingGame, index int) int {
	return MaxScore + game.rolls[index+2]
}

type DefaultCalculator struct{}

func (c DefaultCalculator) shift() int {
	return 2
}

func (c DefaultCalculator) calculate(game BowlingGame, index int) int {
	return game.rolls[index] + game.rolls[index+1]
}

type Score struct {
	score     int
	rollIndex int
}

func (score Score) add(points int, addIndex int) Score {
	return Score{score.score + points, score.rollIndex + addIndex}
}

type BowlingGame struct {
	rolls []int
}

func (game *BowlingGame) Roll(hits int) {
	game.rolls = append(game.rolls, hits)
}

func (game BowlingGame) CalculateScore() int {
	score := Score{0, 0}

	for i := 0; i < 10; i++ {
		calculator := game.getCalculator(score.rollIndex)
		score = score.add(calculator.calculate(game, score.rollIndex), calculator.shift())
	}

	return score.score
}

func (game BowlingGame) getCalculator(index int) ScoreCalculator {
	if game.isStrike(index) {
		return StrikeCalculator{}
	} else if game.isSpare(index) {
		return SpareCalculator{}
	}
	return DefaultCalculator{}
}

func (game BowlingGame) isStrike(index int) bool {
	return game.rolls[index] == MaxScore
}

func (game BowlingGame) isSpare(index int) bool {
	return game.rolls[index]+game.rolls[index+1] == MaxScore
}
