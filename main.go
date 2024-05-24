package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"os"
)

func forwardHandler(w http.ResponseWriter, req *http.Request) {
	baseurl := os.Getenv("HTTP_FORWARD_BASEURL")
	if baseurl == "" {
		log.Fatal("HTTP_FORWARD_BASEURL env variable must be set")
	}

	forward, err := http.NewRequest(req.Method, baseurl+req.RequestURI, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	forward.Header = req.Header

	resp, err := http.DefaultClient.Do(forward)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	port := os.Getenv("HTTP_FORWARD_PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	http.HandleFunc("/", forwardHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

