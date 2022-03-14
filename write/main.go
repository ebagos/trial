package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	WAIT = 5
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(WAIT * time.Second)
	fmt.Fprintf(w, "Hello Write World")
}

func main() {
	myPort := os.Getenv("PORT")
	if myPort == "" {
		myPort = ":8082"
	}
	if strings.HasPrefix(myPort, ":") == false {
		myPort = ":" + myPort
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(myPort, nil)
}
