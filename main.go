package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//define structs

type Movies struct {
	Id       string    `json : "id"`
	Imdb     float32   `json : "imdb"`
	Name     string    `json : "name"`
	Director *Director `json : "director"`
}

type Director struct {
	FirstName string `json : "firstname"`
	LastName  string `json : "lastname"`
	Age       int    `json : "age"`
}

// create movies
var movies []Movies

// create route functions
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type ", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)

	for index, value := range movies {
		if value.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, value := range movies {
		if value.Id == params["id"] {
			json.NewEncoder(w).Encode(value)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index , value := range movies{
		if value.Id == params["id"] {
			var movie Movies
			movie.Id = strconv.Itoa(rand.Intn(100))
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movies = append(movies[:index],movie)
			movies = append(movies, movies[index+2:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	fmt.Println("Welcome to learning crud operation")

	//create handler
	r := mux.NewRouter()

	//create few movies for testing
	movies = append(movies, Movies{Id: "1", Imdb: 9.8, Name: "Lucy", Director: &Director{FirstName: "Luc", LastName: "Bensson", Age: 54}})
	movies = append(movies, Movies{Id: "2", Imdb: 8.8, Name: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan", Age: 64}})

	//create routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting the server\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
