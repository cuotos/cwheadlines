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

	p := problem.GenerateProblem(a.Title)

	fmt.Println(p.AsMorse())

	for {
		userInput := bufio.NewScanner(os.Stdin)
		userInput.Scan()

		if strings.HasPrefix(userInput.Text(), "/") {

			handlers.HandleSlashCmds(userInput.Text(), p)

		} else {
			correct, message := handlers.GuessHandler(a.Title, userInput.Text())

			if correct {
				break
			} else {
				fmt.Println(message)
			}
		}

	}
	fmt.Println("Well Dome")
}
