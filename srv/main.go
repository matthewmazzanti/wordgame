package main

import (
	"io"
	"log"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	fmt.Println("test")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	EnableCors(&w)
	io.WriteString(w, "hello world!")
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
