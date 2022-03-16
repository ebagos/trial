package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var dbURL string = "http://db:8083/write"

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(dbURL)
	if err != nil {
		log.Fatal("write http error :", err)
	}
	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("write readAll: ", err)
	}
	fmt.Fprintf(w, "%s", string(byteArray))
}

func main() {
	myPort := os.Getenv("PORT")
	if myPort == "" {
		myPort = ":8082"
	}
	dburl := os.Getenv("DB_URL")
	if dburl != "" {
		dbURL = dburl
	}
	if !strings.HasPrefix(myPort, ":") {
		myPort = ":" + myPort
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(myPort, nil)
}
