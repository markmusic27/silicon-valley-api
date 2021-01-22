package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Quote string `json:"Quote"`
	SaidBy string `json:"SaidBy"`
	Context string `json:"Context"`
	Episode string `json:"Episode"`
	Season string `json:"Season"`
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}