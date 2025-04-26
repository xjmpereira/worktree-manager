package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	choice         FuzzyMatch
	BorderStyle    lipgloss.Style = lipgloss.NewStyle().
											Background(lipgloss.AdaptiveColor{Light: "247", Dark: "240"})
	normalStyle    lipgloss.Style = lipgloss.NewStyle().
											Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "7"})
	sltStyle       lipgloss.Style = lipgloss.NewStyle().
											Foreground(lipgloss.AdaptiveColor{Light: "236", Dark: "174"}).
											Background(lipgloss.AdaptiveColor{Light: "236", Dark: "236"}).
											Bold(true)
	sltBorderStyle lipgloss.Style = lipgloss.NewStyle().
											Background(lipgloss.AdaptiveColor{Light: "236", Dark: "246"})
)

type (
	errMsg error
)

type FuzzyMatch struct {
	Path     string
    Branch   string
    Distance int
}

type model struct {
	textInput     textinput.Model
	windowWidth   int
	windowHeight  int
	worktreeList  []GitwsWorktree
	rankedList    []FuzzyMatch
	err           error
}

func initialModel(wsList []GitwsWorktree) model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 260
	ti.Width = 20

	return model{
		textInput:    ti,
		worktreeList: wsList,
		err:          nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.rankedList = Map(m.worktreeList, func(a GitwsWorktree) FuzzyMatch {
		percent := 0.8
		distance := SimilarText(a.Branch, m.textInput.Value(), &percent)
		return FuzzyMatch{
			Path: a.Path,
			Branch: a.Branch,
			Distance: distance,
		}
	})
	sort.Slice(m.rankedList, func(i, j int) bool {
		return m.rankedList[i].Distance > m.rankedList[j].Distance
	})
	choice = m.rankedList[0]
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var s strings.Builder
	titleHeight := 0
	footerHeight := 3
	// Worktree list
	maxRow := m.windowHeight-footerHeight-titleHeight
	for row := range maxRow {
		index := maxRow-row-1
		if index < len(m.rankedList) {
			if index > 0 {
				s.WriteString(BorderStyle.Render(" "))
				s.WriteString(normalStyle.Render(" " + m.rankedList[index].Branch) + " ")
			} else if index == 0 {
				s.WriteString(sltBorderStyle.Render(" "))
				s.WriteString(sltStyle.Render(" " + choice.Branch + " "))
			}
		}
		s.WriteString("\n")
	}
	// Footer section
	s.WriteString("\n")
	s.WriteString(m.textInput.View())
	s.WriteString("\n")
	return s.String()
}

func iterative(wsList []GitwsWorktree) FuzzyMatch {
	m := initialModel(wsList)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return choice
}
