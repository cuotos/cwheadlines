package handlers

import "fmt"

func HandleSlashCmds(s string){
	slashCommand := s[1:]

	switch slashCommand {
	case "giveup":
		fmt.Println("giveup")
	default:
	}
}
