package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = ""

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. This is our first Go web app on Heroku! Local port is "+port)
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
