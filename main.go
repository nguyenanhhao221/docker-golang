package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
		}
		if _, err := w.Write([]byte("Hello, Mom")); err != nil {
			log.Printf("Error when write data for root path: %v", err)
			return
		}
	})
	http.HandleFunc("/health", handleCheckHealth)

	PORT := ":8080"

	log.Println("Server is listening on port", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatalln("Error Failed to start server", err)
	}
}

func handleCheckHealth(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}

	payload := Response{
		Status: "ok",
	}

	responseWithJson(w, http.StatusOK, payload)
}

func responseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error failed to marshal data to json", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		log.Println("Error failed to write data", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
