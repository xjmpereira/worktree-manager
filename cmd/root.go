package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	cli "github.com/urfave/cli/v3"
)

var (
	choice         FuzzyMatch
	currentDir     string
	foundConfig    bool = false
	gitwsConfig    GitwsConfig
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

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 260
	ti.Width = 20

	wsList := GetWsList()
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

// SimilarText calculates the similarity between two strings.
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

func RootCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "ws",
		Usage: "A Program to manage git worktrees",
		Before: RootBeforeFn,
		Action: RootFn,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "currentdir",
				Value: "",
				Aliases: []string{"C"},
				Usage: "Use this as the current directory",
			},
		},
		Commands: []*cli.Command{
			ConfigCmd(),
			CloneCmd(),
			ListCmd(),
			CreateCmd(),
		},
	}
	return cmd
}

func RootBeforeFn(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	RmPostCmdFile()
	currentDir = cwdOr(cmd.String("currentdir"))
	rootDir, err := searchGitwsRoot(currentDir)
	if err == nil {
		foundConfig = true
		gitwsConfig = readGitwsConfig(rootDir)
	}
	return nil, nil
}

func RootFn(ctx context.Context, cmd *cli.Command) error {
	m := initialModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	
	SetPostCmd("cd " + choice.Path)
	return nil
}

func RmPostCmdFile() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	
	directory := filepath.Join(dirname, ".config", "gitws")
	path := filepath.Join(directory, "ws-post")
	os.Remove(path)
}

func SetPostCmd(cmd string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	
	directory := filepath.Join(dirname, ".config", "gitws")
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(directory, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}
	
	path := filepath.Join(directory, "ws-post")
	err = os.WriteFile(path, []byte(cmd), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
