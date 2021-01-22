package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Quote string `json:"Quote"`
	SaidBy string `json:"SaidBy"`
	Context string `json:"Context"`
	Episode int `json:"Episode"`
	Season int `json:"Season"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
	Article{Quote:"Test", SaidBy: "Test", Context: "Test", Episode: 1, Season: 6},
	}

	fmt.Println(w, "Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}