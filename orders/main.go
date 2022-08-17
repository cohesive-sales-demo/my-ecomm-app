package main

import (
	"fmt"
	"github.com/testOrgHimali/app/pkg/utils"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprintf(w, "<html> <head></head><body><h1>Welcome to Cohesive</h1> <br> <h2>Hello Orders %d \n </h2></body></html>", utils.Trigger())
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
