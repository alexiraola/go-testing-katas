package main

import (
	"testing"
)

func rollMany(game *BowlingGame, times int, pins int) {
	for i := 0; i < times; i++ {
		game.Roll(pins)
	}
}

func TestShoulBeAbleToRollABall(t *testing.T) {
	game := BowlingGame{}

	game.Roll(0)

	result := game.rolls[0]

	if result != 0 {
		t.Fatalf(`expected %d, got %d`, 0, result)
	}
}

func TestZeroHitsShouldGetZeroPoints(t *testing.T) {
	game := BowlingGame{}

	rollMany(&game, 20, 0)

	result := game.CalculateScore()

	if result != 0 {
		t.Fatalf(`expected %d, got %d`, 0, result)
	}
}

func TestOneHitEveryShotShouldGet20Points(t *testing.T) {
	game := BowlingGame{}

	rollMany(&game, 20, 1)

	result := game.CalculateScore()

	if result != 20 {
		t.Fatalf(`expected %d, got %d`, 20, result)
	}
}
