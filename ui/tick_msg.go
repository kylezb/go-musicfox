package ui

import (
	tea "github.com/anhoder/bubbletea"
	"time"
)

// startup tick
type tickStartupMsg struct{}

func tickStartup(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(time.Time) tea.Msg {
		return tickStartupMsg{}
	})
}

// main ui tick
type tickMainUIMsg struct{}

func tickMainUI(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(time.Time) tea.Msg {
		return tickMainUIMsg{}
	})
}