package main

import (
	"log"
	"net/http"
	"os"
)

type PlaygroundTest struct {
	StratosAuthPort string
	TestThing       uint8
}

type IPlaygroundTest interface {
	DoSomething()
}

func (p *PlaygroundTest) Init() {
	p.StratosAuthPort = "8080"
	p.TestThing = 1
}

func (p *PlaygroundTest) DoSomething() {
	log.Println("Doing something")
}

func PassInterface(p IPlaygroundTest) {
	p.DoSomething()
}

func main() {
	log.Println("Starting the application...")

	// Get the PORT from the environment variables
	port := os.Getenv("STRATOS_AUTH_PORT")
	log.Println("STRATOS_AUTH_PORT: ", port)

	// Create a new PlaygroundTest object
	playgroundTest := PlaygroundTest{}
	playgroundTest.Init()

	PassInterface(&playgroundTest)

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
