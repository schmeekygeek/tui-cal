package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
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

	log.Println("starting ssh server on port 2222...")
	log.Printf("connections will only last %s\n", DeadlineTimeout)
	log.Printf("and timeout after %s of no activity\n", IdleTimeout)
	server := &ssh.Server{
		Addr:        ":2222",
		MaxTimeout:  DeadlineTimeout,
		IdleTimeout: IdleTimeout,
	}
  _ = server
	// log.Fatal()
}
