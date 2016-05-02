package main

import "net/http"

func main() {
	http.HandleFunc("/times", timesHandler)
	http.ListenAndServe(":8080", nil)
}
