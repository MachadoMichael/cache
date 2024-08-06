package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MachadoMichael/cache/infra/database"
	"github.com/MachadoMichael/cache/omdb"
	"github.com/MachadoMichael/cache/schema"
)

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/omdb/{movie}", searchMovie)
	mux.HandleFunc("/cache/{movie}", getCachedMovie)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	srv.ListenAndServe()

}

func searchMovie(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	movie := r.PathValue("movie")

	sr, err := omdb.GetMovies(movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	dataInString, err := json.Marshal(sr.Search)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	dataToCache := schema.DataCache{
		Key:  movie,
		Data: string(dataInString),
	}

	database.CacheRepo.Create(dataToCache)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sr.Search)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func getCachedMovie(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	movie := r.PathValue("movie")

	dataAsString, err := database.CacheRepo.ReadOne(movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var movies []omdb.Movie
	elapsed := time.Since(start)

	json.Unmarshal([]byte(dataAsString), &movies)

	for _, movie := range movies {
		fmt.Println(movie.Title)
	}
	fmt.Println(elapsed)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataAsString)

}
