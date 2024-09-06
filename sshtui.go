package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
  expr string
  res  string
  cursorPos Point
  buttons [][]string
}

type Point struct {
  x, y int
}

func initialModel() model {
  return model{
  	expr:      "",
  	res:       "",
  	cursorPos: Point{x: 0, y: 0},
  	buttons:   [][]string{
      {"1", "2", "3"},
      {"4", "5", "6"},
      {"7", "8", "9"},
      {"0", "-", "*"},
      {"/", "+", "="},
    },
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
      case "ctrl+c", "q":
      return m, tea.Quit

      case "up", "k":
      if m.cursorPos.y != 0 {
        m.cursorPos.y--
      }

      case "down", "j":
      if m.cursorPos.y != len(m.buttons) - 1 {
        m.cursorPos.y++
      }

      case "left", "h":
      if m.cursorPos.x != 0 {
        m.cursorPos.x--
      }

      case "right", "l":
      if m.cursorPos.x != len(m.buttons[0]) - 1 {
        m.cursorPos.x++
      }
    }
  }
  return m, nil
}

func (m model) View() string {
  s := ""
  s += strings.Repeat(" ", m.cursorPos.x * 2)
  s += "."
  s += "\n"
  return s
}
