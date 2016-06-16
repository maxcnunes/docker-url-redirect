package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const port = "8080"

func main() {
	url := os.Getenv("REDIRECT_URL")
	code, err := strconv.Atoi(os.Getenv("REDIRECT_STATUS"))
	if err != nil {
		code = http.StatusTemporaryRedirect
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, code)
	})

	fmt.Printf("Starting server at %s. Redirecting all requests to %s with status %d.\n", port, url, code)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
