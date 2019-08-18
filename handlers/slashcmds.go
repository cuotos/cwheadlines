package handlers

import (
	"cwheadlines/problem"
	"fmt"
)

func HandleSlashCmds(s string, p problem.Problem) {
	slashCommand := s[1:]

	switch slashCommand {
	case "giveup":
		fmt.Println(p.AsString())
	default:
	}
}
