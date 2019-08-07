package guess

import (
	"fmt"
)


func GuessHandler(solution, guess string) (bool, string) {

		switch g := guess; g {
		case solution:
			return true, ""
		case "giveup":
			fmt.Println(solution)
			return false, ""
		default:
			return false, "incorrect. try again"
		}
}
