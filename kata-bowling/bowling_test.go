package main

import (
	"testing"
)

func rollMany(game *BowlingGame, times int, pins int) {
	for i := 0; i < times; i++ {
		game.Roll(pins)
	}
}

func rollManyCouple(game *BowlingGame, times int, first int, second int) {
	for i := 0; i < times; i++ {
		game.Roll(first)
		game.Roll(second)
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

func TestCalculatesTheScoreForAGivenSpareExtraBall(t *testing.T) {
	game := BowlingGame{}

	game.Roll(5)
	game.Roll(5)
	game.Roll(5)
	rollMany(&game, 17, 0)

	result := game.CalculateScore()

	if result != 20 {
		t.Fatalf(`expected %d, got %d`, 20, result)
	}
}

func TestCalculatesTheScoreForAGivenStrikeExtraBall(t *testing.T) {
	game := BowlingGame{}

	game.Roll(10)
	game.Roll(2)
	game.Roll(3)
	rollMany(&game, 16, 0)

	result := game.CalculateScore()

	if result != 20 {
		t.Fatalf(`expected %d, got %d`, 20, result)
	}
}

func TestCalculatesAPerfectGame(t *testing.T) {
	game := BowlingGame{}

	rollMany(&game, 12, 10)

	result := game.CalculateScore()

	if result != 300 {
		t.Fatalf(`expected %d, got %d`, 300, result)
	}
}

func TestCalculatesASpareInTheLastTurn(t *testing.T) {
	game := BowlingGame{}

	rollMany(&game, 21, 5)

	result := game.CalculateScore()

	if result != 150 {
		t.Fatalf(`expected %d, got %d`, 150, result)
	}
}

func TestCalculatesASpareWithDifferentScoreTheLastTurn(t *testing.T) {
	game := BowlingGame{}

	rollManyCouple(&game, 10, 8, 2)
	game.Roll(8)

	result := game.CalculateScore()

	if result != 180 {
		t.Fatalf(`expected %d, got %d`, 180, result)
	}
}
