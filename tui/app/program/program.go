package program

import (
	"fmt"
	"log/slog"

	tea "charm.land/bubbletea/v2"
)

var teaProgram *tea.Program

func GetProgram() *tea.Program {
	return teaProgram
}

func SetProgram(p *tea.Program) {
	teaProgram = p
}

func Send[T tea.Msg](msg T) error {
	if teaProgram == nil {
		slog.Error("program is not set")

		return fmt.Errorf("program is not set")
	}

	go teaProgram.Send(msg)

	return nil
}
