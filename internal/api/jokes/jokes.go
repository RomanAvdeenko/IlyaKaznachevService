package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

func (jc *JokeClient) GetJoke() (string, error) {
	urlPath := jc.url + getJokePath
	resp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	} else if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request error: %v", err)
	}

	defer resp.Body.Close()

	var data JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return "", err
}

type JokeResponse struct {
	Joke string `json:"joke"`
}
