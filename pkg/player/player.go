package player

import (
	"fmt"
	"strings"
)

type Player struct {
	label string
	score float64
}

func NewPlayer(label string, options ...Option) (*Player, error) {
	p := &Player{
		label: label,
		score: 0,
	}

	defaultArgs := &Options{
		InitialScore: 0,
	}

	for _, opt := range options {
		opt(defaultArgs)
	}

	p.score = defaultArgs.InitialScore

	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) Label() string {
	return p.label
}

func (p *Player) Score() float64 {
	return p.score
}

func (p *Player) Validate() error {
	label := strings.ReplaceAll(p.label, " ", "")
	if len(label) == 0 {
		return fmt.Errorf("action=player.validate field=%v err=%w", "label", ErrMissingField)
	}
	return nil
}

func (p *Player) IncrementScore(v float64) {
	p.score += v
}
