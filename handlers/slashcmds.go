package handlers

import (
	"github.com/cuotos/cwheadlines/problem"
	"fmt"
)

func HandleSlashCmds(s string, p problem.Problem) {
	slashCommand := s[1:]

	switch slashCommand {
	case "giveup":
		fmt.Println(p.Solution)
	default:
	}
}
