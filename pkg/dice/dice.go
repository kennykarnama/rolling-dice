package dice

import (
	"rolling-dice/util/random"
)

type Dice struct {
	Values []int
}

func NewDice(n int, options ...Option) *Dice {
	d := &Dice{
		Values: []int{
			0,
			0,
			0,
			0,
			0,
			0,
		},
	}

	defaultArgs := &Options{
		RandomValueGen: random.RandIntInterval,
	}

	for _, opt := range options {
		opt(defaultArgs)
	}

	randVal := defaultArgs.RandomValueGen(1, 6, n)
	for idx := 0; idx < n; idx++ {
		d.Values[randVal[idx]-1]++
	}

	return d
}

func (d *Dice) RemoveValue(v int, freq int) {
	d.Values[v-1] -= freq
}

func (d *Dice) GetValue(v int) int {
	return d.Values[v-1]
}

func (d *Dice) IncrementValue(v int) {
	d.Values[v-1]++
}
