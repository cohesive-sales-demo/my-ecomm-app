package main

import (
	"fmt"
	"github.com/testOrgHimali/app/pkg/utils"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprintf(w, "<html> <head></head><body><h1>Hello Orders %d \n </h1></body></html>", utils.Trigger())
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
