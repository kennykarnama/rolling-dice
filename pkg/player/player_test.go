package player_test

import (
	"errors"
	"rolling-dice/pkg/player"
	"testing"
)

func TestNewPlayer__WithOutError(t *testing.T) {
	args := [][]interface{}{
		[]interface{}{
			"A",
		},
		[]interface{}{
			"A",
			float64(10),
		},
	}

	for _, arg := range args {
		var p *player.Player
		var err error
		expectedLabel := arg[0].(string)
		var expectedScore float64
		if len(arg) == 1 {
			p, err = player.NewPlayer(expectedLabel)
			if err != nil {
				t.Errorf("found err: %v", err)
				return
			}
		} else if len(arg) == 2 {
			expectedScore = arg[1].(float64)
			p, err = player.NewPlayer(expectedLabel, player.SetInitialScore(expectedScore))
			if err != nil {
				t.Errorf("found err: %v", err)
				return
			}

		} else {
			t.Errorf("number of args exceeds. found: %v", len(arg))
		}

		if p == nil {
			t.Errorf("expected p obj not nil")
			return
		}
		if p.Label() != expectedLabel {
			t.Errorf("expected label: %v but found: %v", expectedLabel, p.Label())
			return
		}
		if p.Score() != expectedScore {
			t.Errorf("expected p.score: %v but found: %v", expectedScore, p.Score())
			return
		}
	}
}

func TestNewPlayerWithErrValidation(t *testing.T) {
	var testcases = []struct {
		In          string
		ErrInstance error
	}{
		{
			In:          "",
			ErrInstance: player.ErrMissingField,
		},
		{
			In:          "     ",
			ErrInstance: player.ErrMissingField,
		},
		{
			In:          "         kenny",
			ErrInstance: nil,
		},
	}

	for _, tt := range testcases {
		_, err := player.NewPlayer(tt.In)
		if tt.ErrInstance != nil {
			if !errors.Is(err, tt.ErrInstance) {
				t.Errorf("expected error of type: %v but found: %v", player.ErrMissingField, err)
				return
			}
		} else {
			if err != nil {
				t.Errorf("expected err nil but found: %v", err)
				return
			}
		}
	}
}
