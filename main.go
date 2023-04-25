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

type VideoGame struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Genre     string     `json:"genre"`
	Developer *Developer `json:"developer & publisher"`
}

type Developer struct {
	DeveloperName string `json:"developer_name"`
	PublisherName string `json:"publisher_name"`
}

var games []VideoGame

func getVideoGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func deleteVideoGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, game := range games {
		if game.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(games)
}

func getVideoGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, game := range games {
		if game.ID == params["id"] {
			json.NewEncoder(w).Encode(game)
		}
	}
}

func createVideoGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var game VideoGame
	_ = json.NewDecoder(r.Body).Decode(&game)
	game.ID = strconv.Itoa(rand.Intn(10000))
	games = append(games, game)
	json.NewEncoder(w).Encode(game)
}

func updateVideoGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, game := range games {
		// deleting the video game if it exists in the slice
		if game.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)
			// Add updated game body into the slice
			var game VideoGame
			_ = json.NewDecoder(r.Body).Decode(&game)
			game.ID = params["id"]
			games = append(games, game)
			json.NewEncoder(w).Encode(game)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	games = append(games, VideoGame{ID: "1", Name: "Destiny 2", Genre: "FPS/RPG", Developer: &Developer{DeveloperName: "Bungie", PublisherName: "Bungie"}})
	games = append(games, VideoGame{ID: "2", Name: "Call of Duty: Black ops 3", Genre: "FPS", Developer: &Developer{DeveloperName: "Activision", PublisherName: "Treyarch"}})
	games = append(games, VideoGame{ID: "3", Name: "Witcher 3", Genre: "RPG", Developer: &Developer{DeveloperName: "CD Projekt RED", PublisherName: "CD Project"}})

	r.HandleFunc("/videogames", getVideoGames).Methods("GET")
	r.HandleFunc("/videogames/{id}", getVideoGame).Methods("GET")
	r.HandleFunc("/videogames", createVideoGame).Methods("POST")
	r.HandleFunc("/videogames/{id}", updateVideoGame).Methods("PUT")
	r.HandleFunc("/videogames/{id}", deleteVideoGame).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
