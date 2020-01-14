package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "Hallo")
	fmt.Printf("%i", counter)
	mutex.Unlock()
}

func main() {

	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi big luck")
	})

	http.HandleFunc("/frode", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "frode is here")
	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
