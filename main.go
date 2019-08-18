package main

import (
	"bufio"
	"cwheadlines/handlers"
	"cwheadlines/problem"
	"cwheadlines/retreive"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a, err := retreive.GetTopStory()
	if err != nil {
		log.Fatal(err)
	}

	normalizedTitle := normalizeText(a.Title)

	p := problem.GenerateProblem(normalizedTitle)

	fmt.Println(p.AsMorse())

	for {
		userInput := bufio.NewScanner(os.Stdin)
		userInput.Scan()

		if strings.HasPrefix(userInput.Text(), "/") {

			handlers.HandleSlashCmds(userInput.Text(), p)

		} else {
			correct, message := handlers.GuessHandler(normalizedTitle, userInput.Text())

			if correct {
				break
			} else {
				fmt.Println(message)
			}
		}

	}
	fmt.Println("Well Dome")
}

func normalizeText(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal("normalization regex failed")
	}
	return strings.ToLower(reg.ReplaceAllString(s, ""))
}
