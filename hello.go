package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	card := model.Card{ID: 1, Title: "Hello"}
	json.NewEncoder(w).Encode(card)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	card := model.Card{ID: 1, Title: "Hello"}
	json.NewEncoder(w).Encode(card)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
