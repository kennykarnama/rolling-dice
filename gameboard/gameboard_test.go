package gameboard_test

import (
	"fmt"
	"rolling-dice/gameboard"
	"rolling-dice/pkg/dice"
	"testing"
)

func TestGameBoard__Evaluate(t *testing.T) {
	gb, err := gameboard.NewGameBoard(3, 4)
	if err != nil {
		t.Errorf("found error: %v", err)
	}
	// p1
	// 3,6,1,3
	func1 := dice.SetRandomValueGen(func(min int, max int, num int) []int {
		return []int{
			3,
			6,
			1,
			3,
		}
	})
	gb.Players[0].Dice = dice.NewDice(4, func1)

	// p2
	// 2,4,5,5
	func2 := dice.SetRandomValueGen(func(min int, max int, num int) []int {
		return []int{
			2,
			4,
			5,
			5,
		}
	})
	gb.Players[1].Dice = dice.NewDice(4, func2)

	// p3
	// 1,2,5,6
	func3 := dice.SetRandomValueGen(func(min int, max int, num int) []int {
		return []int{
			1,
			2,
			5,
			6,
		}
	})
	gb.Players[2].Dice = dice.NewDice(4, func3)

	gb.Evaluate()

	expectedNumDice := []int{
		3,
		5,
		2,
	}

	expectedScore := []float64{
		1,
		0,
		1,
	}

	for idx, player := range gb.Players {

		if player.CurrentNumDice != expectedNumDice[idx] {
			t.Errorf("player: %v expected numDice: %v but found: %v", player.Player.Label(), expectedNumDice[idx], player.CurrentNumDice)
			return
		}

		if player.Player.Score() != expectedScore[idx] {
			t.Errorf("expected numDice: %v but found: %v", expectedScore[idx], player.Player.Score())
			return
		}
	}
}

func TestGameBoard__EvaluateStop(t *testing.T) {
	gb, err := gameboard.NewGameBoard(3, 4)
	if err != nil {
		t.Errorf("found error: %v", err)
	}
	// p1
	// 6,1
	func1 := dice.SetRandomValueGen(func(min int, max int, num int) []int {
		return []int{
			6,
			1,
		}
	})
	gb.Players[0].Dice = dice.NewDice(2, func1)
	gb.Players[0].CurrentNumDice = 2
	gb.Players[0].Player.IncrementScore(2)
	fmt.Println(gb.Players[0].Dice)

	// p2
	// 2,5,6,4,6
	func2 := dice.SetRandomValueGen(func(min int, max int, num int) []int {
		return []int{
			2,
			5,
			6,
			4,
			6,
		}
	})
	gb.Players[1].Dice = dice.NewDice(5, func2)
	gb.Players[1].CurrentNumDice = 5

	// p3
	// 1
	func3 := dice.SetRandomValueGen(func(min int, max int, num int) []int {
		return []int{
			1,
		}
	})
	gb.Players[2].Dice = dice.NewDice(1, func3)
	gb.Players[2].CurrentNumDice = 1
	gb.Players[2].Player.IncrementScore(2)

	gb.Evaluate()

	expectedNumDice := []int{
		1,
		4,
		0,
	}

	expectedScore := []float64{
		3,
		2,
		2,
	}

	for idx, player := range gb.Players {

		if player.CurrentNumDice != expectedNumDice[idx] {
			t.Errorf("player: %v expected numDice: %v but found: %v", player.Player.Label(), expectedNumDice[idx], player.CurrentNumDice)
			return
		}

		if player.Player.Score() != expectedScore[idx] {
			t.Errorf("player: %v expected score: %v but found: %v", player.Player.Label(), expectedScore[idx], player.Player.Score())
			return
		}
	}

	// player 3 should exit

	fmt.Println(gb.Players[2].CurrentNumDice)
	if !gb.Players[2].Exit {
		t.Errorf("expected player 3 to exit")
		return
	}
}

func TestWinner(t *testing.T) {
	gb, err := gameboard.NewGameBoard(3, 4)
	if err != nil {
		t.Errorf("found error: %v", err)
	}
	gb.Players[0].Player.IncrementScore(3)
	gb.Players[1].Player.IncrementScore(2)
	gb.Players[2].Player.IncrementScore(2)
	winner := gb.Winner()
	if winner.Player.Label() != gb.Players[0].Player.Label() {
		t.Errorf("expected winner is: %v but found: %v", gb.Players[0].Player.Label(), winner.Player.Label())
		return
	}
}
func TestPlay(t *testing.T) {
	gb, err := gameboard.NewGameBoard(3, 4)
	if err != nil {
		t.Errorf("found error: %v", err)
	}
	winner, err := gb.Play()
	if err != nil {
		t.Errorf("found error: %v", err)
		return
	}
	fmt.Println(winner.Player.Label())
}
