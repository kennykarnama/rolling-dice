package gameboard

import (
	"fmt"
	"rolling-dice/pkg/dice"
	"rolling-dice/pkg/player"
	"sort"
)

type GameBoardPlayer struct {
	Player         *player.Player
	Dice           *dice.Dice
	CurrentNumDice int
	Exit           bool
}

type GameBoard struct {
	currentNumPlayer int
	Players          []*GameBoardPlayer
}

func NewGameBoard(numPlayer int, numDice int) (*GameBoard, error) {

	gbPlayers := []*GameBoardPlayer{}
	for pIdx := 1; pIdx <= numPlayer; pIdx++ {
		p, err := player.NewPlayer(fmt.Sprintf("player %d", pIdx))
		if err != nil {
			return nil, err
		}
		gbPlayer := &GameBoardPlayer{
			Player:         p,
			CurrentNumDice: numDice,
		}
		gbPlayers = append(gbPlayers, gbPlayer)

	}

	return &GameBoard{
		currentNumPlayer: numPlayer,
		Players:          gbPlayers,
	}, nil
}

func (gb *GameBoard) Play() (winner *GameBoardPlayer, err error) {
	for !gb.playStop() {

		// roll the dice
		gb.RollDice()

		// evaluate
		gb.Evaluate()
	}

	return gb.Winner(), nil
}

func (gb *GameBoard) RollDice(option ...dice.Option) {
	for _, player := range gb.Players {
		if player.Exit {
			continue
		}
		player.Dice = dice.NewDice(player.CurrentNumDice, option...)
	}
}

func (gb *GameBoard) Evaluate() {
	for curIdx, player := range gb.Players {
		if player.Exit {
			//log.Printf("player=%v has exitted", player.Player.Label())
			continue
		}
		// check if has 6
		if v := player.Dice.GetValue(6); v > 0 {
			player.Dice.RemoveValue(6, v)
			player.CurrentNumDice -= v
			player.Player.IncrementScore(float64(v))
		}

		// give to the next
		if v := player.Dice.GetValue(1); v > 0 {
			cur := curIdx
			nextIdx := (cur + 1) % len(gb.Players)
			for nextIdx != cur {
				if !gb.Players[nextIdx].Exit {
					gb.Players[nextIdx].CurrentNumDice = gb.Players[nextIdx].CurrentNumDice + 1
					player.CurrentNumDice -= v
					player.Dice.RemoveValue(1, v)
					break
				}
				nextIdx = (nextIdx + 1) % len(gb.Players)
			}
		}
	}

	// mark exit if player numDice == 0
	for _, player := range gb.Players {
		if player.CurrentNumDice == 0 {
			player.Exit = true
			gb.currentNumPlayer--
		}
	}
}

func (gb *GameBoard) playStop() bool {
	return gb.currentNumPlayer == 1
}

func (gb *GameBoard) Winner() *GameBoardPlayer {
	sort.Slice(gb.Players[:], func(i, j int) bool {
		return gb.Players[i].Player.Score() > gb.Players[j].Player.Score()
	})
	return gb.Players[0]
}
