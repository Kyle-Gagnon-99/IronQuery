package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting the application...")

	// Get the PORT from the environment variables
	port := os.Getenv("STRATOS_AUTH_PORT")
	log.Println("STRATOS_AUTH_PORT: ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received a request")
		log.Println("Request URL: ", r.URL)
		w.Write([]byte("Hello, World!"))
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
