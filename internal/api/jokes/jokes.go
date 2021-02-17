package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"workshop/internal/api"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

func NewJokeClient(url string) *JokeClient {
	return &JokeClient{
		url: url,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %v", http.StatusText(resp.StatusCode))
	}

	defer resp.Body.Close()

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}
