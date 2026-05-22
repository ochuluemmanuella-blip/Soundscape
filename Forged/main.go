package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("BeatForge running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/beat/generate", func(w http.ResponseWriter, r*http.Request)) {
		fmt.Fprintln(w, "a beat will go here soon")
	}
}
