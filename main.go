package main

import (
	"bufio"
	"cwheadlines/guess"
	"cwheadlines/models"
	"cwheadlines/morse"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
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

	a, err := getTopStory()
	if err != nil {
		log.Fatal(err)
	}

	normalizedTitle := normalizeText(a.Title)


	morseString := convertToMorse(normalizedTitle)
	fmt.Println(morseString)



	for {
		userInput := bufio.NewScanner(os.Stdin)
		userInput.Scan()

		correct, message := guess.GuessHandler(normalizedTitle, userInput.Text())
		if correct {
			break
		} else {
			fmt.Println(message)
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

func convertToMorse(input string) string {
	return morse.AsMorse(input)
}

func getTopStory() (models.Article, error) {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		return models.Article{}, errors.New("unable to load API_KEY from environment variables")
	}

	url := `https://newsapi.org/v2/top-headlines?sources=bbc-news&apiKey=` + apiKey

	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Article{}, err
	}

	res, err := c.Do(req)
	if err != nil {
		return models.Article{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.Article{}, err
	}

	topStories := &models.TopHeadlinesResponse{}

	err = json.Unmarshal(body, topStories)
	if err != nil {
		return models.Article{}, err
	}

	if topStories.Status == "ok" {
		return topStories.Articles[rand.Intn(len(topStories.Articles))], nil
	}

	return topStories.Articles[0], nil
}
