package retriever

import (
	"github.com/cuotos/cwheadlines/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

type HeadlineRetriever struct{}

func (HeadlineRetriever) GetSolution() (string, error) {
	story, err := GetTopStory()
	if err != nil {
		return "", err
	}

	return story.Title, nil
}

func GetTopStory() (models.Article, error) {
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
