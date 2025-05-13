package models

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ltfred/chronos/internal"
	"strings"
)

type ConfigModel struct {
	keymap   configKeymap
	help     help.Model
	textarea textarea.Model
}

type configKeymap struct {
	saveKeymap
	quitKeymap
}

func NewConfigModel(cfg string) ConfigModel {
	ti := textarea.New()
	ti.Focus()
	ti.SetHeight(strings.Count(cfg, "\n"))
	ti.SetValue(cfg)
	ti.FocusedStyle.CursorLine = cursorLineStyle
	ti.FocusedStyle.Base = textAreaFocusStyle
	k := configKeymap{
		saveKeymap: newSaveKeymap(),
		quitKeymap: newQuitKeymap(),
	}

	return ConfigModel{
		textarea: ti,
		keymap:   k,
		help:     help.New(),
	}
}

func (m ConfigModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m ConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.save):
			internal.SaveConfig(m.textarea.Value())
			m.textarea.Blur()
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m ConfigModel) View() string {
	helpMsg := m.help.ShortHelpView([]key.Binding{
		m.keymap.save,
		m.keymap.quit,
	})

	return fmt.Sprintf(
		"Edit your configuration.\n\n%s",
		m.textarea.View(),
	) + "\n\n" + helpMsg
}
