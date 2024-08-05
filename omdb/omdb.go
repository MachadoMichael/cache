package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SearchResponse struct {
	Search   []Movie `json:"Search"`
	Total    string  `json:"totalResults"`
	Response string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func GetMovies(movie string) (SearchResponse, error) {
	apiKey := "49288edd"
	page := 1
	var searchResponse SearchResponse
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", apiKey, movie, page)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return searchResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return searchResponse, err
	}

	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return searchResponse, err
	}

	fmt.Printf("Total results: %s\n", searchResponse.Total)
	for _, movie := range searchResponse.Search {
		fmt.Printf("Title: %s, Year: %s, Type: %s\n", movie.Title, movie.Year, movie.Type)
	}

	return searchResponse, nil
}
