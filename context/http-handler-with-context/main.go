package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Started")
	defer log.Println("Request Ended")
	w.Header().Set("Content-Type", "application/json")
	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processed successfully")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Request ended successfully"}`))
	case <-ctx.Done():
		log.Println("Request cancelled by client")
	}
}
