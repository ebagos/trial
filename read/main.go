package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Read World")
}

func main() {
	myPort := os.Getenv("PORT")
	if myPort == "" {
		myPort = ":8081"
	}
	if strings.HasPrefix(myPort, ":") == false {
		myPort = ":" + myPort
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(myPort, nil)
}
