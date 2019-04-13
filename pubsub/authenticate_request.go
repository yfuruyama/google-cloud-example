package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		log.Printf(authHeader)
		fmt.Fprintf(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
