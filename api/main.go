package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequest(){
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}