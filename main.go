package main

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"
	"time"
)

func hashHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))

	id := uuid.New()
	w.Write([]byte("\nGenerated UUID: " + id.String()))
}

func main() {

	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// goroutine to launch a server on port 9000
	go func() {
		mux := http.NewServeMux()
		// Convert the timeHandler function to a http.HandlerFunc type.
		th := http.HandlerFunc(hashHandler)
		// And add it to the ServeMux.
		mux.Handle("/hash", th)
		log.Print("Listening(:9000)...")
		log.Fatal(http.ListenAndServe(":9000", mux))
		wg.Done() // one goroutine finished
	}()

	// goroutine to launch a server on port 9001
	go func() {
		mux := http.NewServeMux()
		// Convert the timeHandler function to a http.HandlerFunc type.
		th := http.HandlerFunc(hashHandler)
		// And add it to the ServeMux.
		mux.Handle("/hash", th)
		log.Print("Listening(:9001)...")
		log.Fatal(http.ListenAndServe(":9001", mux))
		wg.Done() // one goroutine finished
	}()

	// wait until WaitGroup is done
	wg.Wait()
}
