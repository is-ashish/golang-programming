package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to golang series of router")
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeRoute).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func welcomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang Projects :)</h1>"))
}
