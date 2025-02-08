package main

import (
	"fmt"
	"os"

	spinner "github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type rootScreenModel struct {
	model tea.Model
}

func RootScreen() rootScreenModel {
	var rootModel tea.Model

	screen_one := screenOne()
	rootModel = &screen_one

	return rootScreenModel{model: rootModel}
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.model.View()
}

func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init()
}

type screenOneModel struct {
	spinner spinner.Model
	err     error
}

type screenTwoModel struct {
	spinner spinner.Model
	err     error
}

func (m screenOneModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m screenOneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			screen_two := screenTwo()
			return screen_two, screen_two.Init()
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

func (m screenOneModel) View() string {
	str := fmt.Sprintf("\n   %s This is screen one...\n\n", m.spinner.View())
	return str
}

func screenTwo() screenTwoModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return screenTwoModel{
		spinner: s,
	}
}

func screenOne() screenOneModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return screenOneModel{
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
			screen_one := screenOne()
			return screen_one, screen_one.Init()
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
	p := tea.NewProgram(RootScreen(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		os.Exit(1)
	}
}
