package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/is-ashish/mongo-api/router"
)

func main() {
	fmt.Println("MongoDB Api")
	r := router.Router()
	fmt.Println("Server Is Getting Started ....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
