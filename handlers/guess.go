package handlers

func GuessHandler(solution, guess string) (bool, string) {

	switch g := guess; g {
	case solution:
		return true, ""
	default:
		return false, "incorrect. try again"
	}
}
