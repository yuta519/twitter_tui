package models

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Accounts     []string
	Choice       int
	IsChosen     bool
	Cursor       int
	Selected     map[int]struct{}
	IsTextFormat bool
	TextInput    textinput.Model
	Tweets       []anaconda.Tweet
	Status       int
	Err          error
}

func Initialize() Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		Accounts:     []string{"golangch", "GolangTrends", "golang_news"},
		IsChosen:     false,
		IsTextFormat: false,
		TextInput:    ti,
		Selected:     make(map[int]struct{}),
	}
}
