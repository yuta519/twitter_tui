package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yuta519/twitter_tui/pkg/tui/models"
)

type model models.Model

func initialize() model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		Accounts:     []string{"golangch", "GolangTrends", "golang_news"},
		IsChosen:     false,
		IsTextFormat: false,
		TextInput:    ti,
		Selected:     make(map[int]struct{}),
	}
}

func main() {
	// program := tea.NewProgram(Init.Initialize())
	program := tea.NewProgram(initialize())
	if err := program.Start(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		// The actual key pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Accounts)-1 {
				m.Cursor++
			}
		case "!":
			m.IsTextFormat = true
			return m, nil
		case " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
				m.IsChosen = false
			} else {
				m.Selected[m.Cursor] = struct{}{}
				m.IsChosen = true
			}
		case "enter":
			anaconda.SetConsumerKey(os.Getenv("API_KEY"))
			anaconda.SetConsumerSecret(os.Getenv("API_SECRET"))
			api := anaconda.NewTwitterApi(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"))
			values := url.Values{}

			if m.IsChosen {
				fmt.Println(m.Accounts[m.Cursor])
				values.Set("screen_name", m.Accounts[m.Cursor])
				m.fetchTweetsByAccount(api, values)
			}
			if m.IsTextFormat {
				values.Set("screen_name", m.TextInput.Value())
				m.fetchTweetsByAccount(api, values)
			}
		}
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
		m.TextInput, cmd = m.TextInput.Update(msg)
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	return m, cmd
}

func (m model) View() string {
	if m.IsTextFormat {
		return textInputView(m)
	}
	return choicesView(m)
}

func choicesView(m model) string {
	// The header
	s := "Choose a twitter account you want check.\n\n"
	// Iterate over our choices
	for i, choice := range m.Accounts {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.Cursor == i {
			cursor = ">" // cursor!
		}
		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.Selected[i]; ok {
			checked = "x" // selected!
		}
		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	// The footer
	s += "\nOr you could input a twitter account with free text when press `!`.\n"
	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}

func textInputView(m model) string {
	return fmt.Sprintf(
		"What’s the twitter account you want know?\n\n%s\n\n%s",
		m.TextInput.View(),
		"(esc to quit)",
	) + "\n"
}

func (m model) fetchTweetsByAccount(api *anaconda.TwitterApi, v url.Values) tea.Msg {
	tweets, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		fmt.Println("tweet: ", tweet.Text)
	}
	return models.Model{Tweets: tweets}
}
