//go:build windows
// +build windows

package tea

import (
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// listenForResize is not available on windows because windows does not
// implement syscall.SIGWINCH.
func (p *Program) listenForResize(done chan struct{}) {
	close(done)
	ticker := time.NewTicker(time.Second)
	var width, height int
	for range ticker.C {
		w, h, err := terminal.GetSize(int(p.output.TTY().Fd()))
		if err != nil {
			close(done)
		}
		if w != width || h != height {
			width, height = w, h
			p.Send(WindowSizeMsg{w, h})
		}
	}

}
