package models

import "github.com/charmbracelet/bubbles/key"

type moveKeymap struct {
	up, down, left, right key.Binding
}

func newMoveKeymap() moveKeymap {
	return moveKeymap{
		up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "left"),
		),
		right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "right"),
		),
	}
}

type quitKeymap struct {
	quit key.Binding
}

func newQuitKeymap() quitKeymap {
	return quitKeymap{
		quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
}

type saveKeymap struct {
	save key.Binding
}

func newSaveKeymap() saveKeymap {
	return saveKeymap{
		save: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "save"),
		),
	}
}

type enterKeymap struct {
	enter key.Binding
}

func newEnterKeymap() enterKeymap {
	return enterKeymap{
		enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter"),
		),
	}
}
