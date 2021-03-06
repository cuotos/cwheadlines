package main

import (
	"bufio"
	"github.com/cuotos/cwheadlines/handlers"
	"github.com/cuotos/cwheadlines/problem"
	"github.com/cuotos/cwheadlines/retriever"
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

	r := retriever.BbcRss{}

	s, err := r.GetSolution()
	if err != nil {
		log.Fatal(err)
	}

	p := problem.GenerateProblem(s)

	fmt.Println(p.AsMorse())

	for {
		userInput := bufio.NewScanner(os.Stdin)
		userInput.Scan()

		if strings.HasPrefix(userInput.Text(), "/") {

			handlers.HandleSlashCmds(userInput.Text(), p)

		} else {
			correct, message := handlers.GuessHandler(p.Solution, userInput.Text())

			if correct {
				break
			} else {
				fmt.Println(message)
			}
		}

	}
	fmt.Println("Well Done")
}
