package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
  Id string `json: "id"`
  Title string `json: "title"`
  Director *Director `json: "director"`
}

type Director struct {
  FirstName string `json: "first_name"`
  SurName string `json: "second_name"`
}

var MovieList []Movie

func deleteMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
    var movies []Movie
    var deletedmovie Movie
  for i, movie := range MovieList {

    if movie.Id == params["id"] {
      movies = append(MovieList[:i], MovieList[i+1:]...) 
      deletedmovie = MovieList[i]
      break
    }
  }
     MovieList = movies
    json.NewEncoder(w).Encode(deletedmovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var updatedmovie Movie
  params := mux.Vars(r)
  json.NewDecoder(r.Body).Decode(&updatedmovie)

  for i, movie := range MovieList {
    if movie.Id == params["id"] {
      movies = append(MovieList[:i], updatedmovie, MovieList[i+1:])
      json.NewEncoder(w).Encode(updatedmovie)
      return
    }
  }
}
func getAllMovies(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(MovieList)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, movie := range MovieList {

    if movie.Id == params["id"] {
      json.NewEncoder(w).Encode(movie)
      return
   }
  }
}
func createMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var movie Movie
  _ = json.NewDecoder(r.Body).Decode(&movie)
  movie.Id = strconv.Itoa(len(MovieList))

  MovieList = append(MovieList, movie)
  json.NewEncoder(w).Encode({
    type: "successfull"
  })
  return
}

func sampleMovies(){
  MovieList = append(MovieList, Movie{Id: "1", Title: "Black Panther", Director: &Director{FirstName: "Amar", SurName: "Jay"}})
  MovieList = append(MovieList, Movie{Id: "2", Title: "Another Movie", Director: &Director{FirstName: "Hello", SurName: "World"}})
}
func main(){
  r := mux.NewRouter()

  sampleMovies()
  r.HandleFunc("/", getAllMovies).Methods("GET")
  r.HandleFunc("/{id}", getMovie).Methods("GET")
  r.HandleFunc("/create/{id}", createMovie).Methods("POST")
  r.HandleFunc("/update/{id}", updateMovie).Methods("PUT")
  r.HandleFunc("/delete/{id}", deleteMovie).Methods("DELETE")

  fmt.Println("Starting server at http://localhost:4000")
  log.Fatal(http.ListenAndServe(":4000", r))
}

