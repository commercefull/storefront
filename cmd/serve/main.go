package main

import (
	"io"
	"log"
	"net/http"
	"storefront/pkg/request"
	"storefront/routes"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	response := routes.Handle(
		request.Request{
			Body:   string(body),
			Path:   r.URL.Path,
			Method: r.Method,
		},
	)

	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", response.Headers["Content-Type"])
	w.Write([]byte(response.Body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})

	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	mux.Handle("/static/", fileHandler)

	log.Print("Starting server on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
