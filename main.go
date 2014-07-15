package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	page, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(rw, string(page))
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
