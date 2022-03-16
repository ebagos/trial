package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func readHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Read World")
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Write World")
}

func main() {
	myPort := os.Getenv("PORT")
	if myPort == "" {
		myPort = ":8083"
	}
	if !strings.HasPrefix(myPort, ":") {
		myPort = ":" + myPort
	}
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(myPort, nil)
}
