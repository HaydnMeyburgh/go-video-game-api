package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type VideoGame struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Developer *Developer `json:"publisher"`
}

type Developer struct {
	developerName string `json:"developername"`
	publisherName string `json:"publishername"`
}

var games []VideoGame

func getVideoGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}



func main() {
	r := mux.NewRouter()

	games = append(games, VideoGame{ID: "1", Name: "Destiny 2", Genre: "FPS/RPG", Developer: &Developer{developerName: "Bungie", publisherName: "Bungie"}})
	games = append(games, VideoGame{ID: "2", Name: "Call of Duty: Black ops 3", Genre: "FPS", Developer: &Developer{developerName: "Activision", publisherName: "Treyarch"}})
	games = append(games, VideoGame{ID: "1", Name: "Witcher 3", Genre: "RPG", Developer: &Developer{developerName: "CD Projekt RED", publisherName: "CD Project"}})
	
	r.HandleFunc("/videogames", getVideoGames).Methods("GET")
	r.HandleFunc("/videogames/{id}", getVideoGame).Methods("GET")
	r.HandleFunc("/videogames", createVideoGame).Methods("POST")
	r.HandleFunc("/videogames/{id}", updateVideoGame).Methods("PUT")
	r.HandleFunc("/videogames/{id}", deleteVideoGame).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}