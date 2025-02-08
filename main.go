package main

import (
	"fmt"
	"os"

	spinner "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type screenOneModel struct {
	title   string
	spinner spinner.Model
	err     error
}

type screenTwoModel struct {
	title   string
	spinner spinner.Model
	err     error
}

func (m screenOneModel) Init() tea.Cmd {
	return nil
}

func (m screenOneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			return screenTwo(), screenTwo().Init()
		}
	}
	return m, nil
}

func (m screenOneModel) View() string {
	str := "This is the first screen. Press any key to switch to the second screen."
	return str
}

func screenTwo() screenTwoModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return screenTwoModel{
		title:   "Loading...",
		spinner: s,
	}
}

func screenOne() screenOneModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return screenOneModel{
		title:   "Loading...",
		spinner: s,
	}
}

func (m screenTwoModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m screenTwoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			return screenOne(), nil
		}
	case error:
		m.err = msg
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m screenTwoModel) View() string {
	str := fmt.Sprintf("\n   %s This is screen two...\n\n", m.spinner.View())
	return str
}

func main() {
	p := tea.NewProgram(screenOne(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		os.Exit(1)
	}
}
