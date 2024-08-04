package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func GetMovies(movie string) {
	start := time.Now()
	apiKey := "49288edd"
	page := 1

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", apiKey, movie, page)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var searchResponse SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("Total results: %s\n", searchResponse.Total)
	for _, movie := range searchResponse.Search {
		fmt.Printf("Title: %s, Year: %s, Type: %s\n", movie.Title, movie.Year, movie.Type)
	}

	elapse := time.Since(start)
	fmt.Println("spent time: ", elapse)
}
