package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	version := "1.0"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	response := fmt.Sprintf("Hello World! API version: %s, Current time: %s", version, currentTime)
	fmt.Fprintln(w, response)
}
