package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"rolling-dice/gameboard"
)

type AppArgs struct {
	NumPlayer int `arg:"--num-player,-n,required" help:"number of player"`
	NumDice   int `arg:"--num-dice,-d,required" help:"number of dice"`
}

var (
	args AppArgs
)

func main() {
	arg.MustParse(&args)
	gb, err := gameboard.NewGameBoard(3, 4)
	if err != nil {
		panic(err)
	}
	winner, err := gb.Play()
	if err != nil {
		panic(err)
	}
	fmt.Println("The winner is", winner.Player.Label())
}
