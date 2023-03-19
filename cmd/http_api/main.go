package main

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func hashHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))

	id := uuid.New()
	w.Write([]byte("\nGenerated UUID: " + id.String()))
}

func main() {
	log.Println("Starting http api server")

	// goroutine to launch a server on port 9000
	mux := http.NewServeMux()
	// Convert the timeHandler function to a http.HandlerFunc type.
	th := http.HandlerFunc(hashHandler)
	// And add it to the ServeMux.
	mux.Handle("/hash", th)
	log.Print("Listening 9000 port...")
	log.Fatal(http.ListenAndServe(":9000", mux))
}
