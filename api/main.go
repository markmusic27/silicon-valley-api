package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Quote struct {
	Quote string `json:"Quote"`
	SaidBy string `json:"SaidBy"`
	Context string `json:"Context"`
	Episode int `json:"Episode"`
	Season int `json:"Season"`
}

type Quotes []Quote

func allQuotes(w http.ResponseWriter, r *http.Request){
	quotes := Quotes{
	Quote{Quote:"Test", SaidBy: "Test", Context: "Test", Episode: 1, Season: 6},
	}

	fmt.Println(w, "Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(quotes)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Silicon Valley API! Here is a list of all of the endpoints:\n\n- ⚡️ All Quotes: /quotes\n- 🧩 Random Quote: /randquote\n\nCheckout the source code at https://github.com/markmusic2727/silicon-valley_api. All contributions are welcome!")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/quotes", allQuotes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}