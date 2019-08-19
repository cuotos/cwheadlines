package handlers

import (
	"cwheadlines/utils"
)

func GuessHandler(solution, guess string) (bool, string) {

	guess = utils.NormalizeText(guess)
	//solution = utils.NormalizeText(solution)

	switch g := guess; g {
	case solution:
		return true, ""
	default:
		return false, "incorrect. try again"
	}
}
