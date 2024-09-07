package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	DeadlineTimeout = 20 * time.Second
	IdleTimeout     = 10 * time.Second
)

func main() {
  p := tea.NewProgram(initialModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("sad %v", err)
    os.Exit(1)
  }
}
