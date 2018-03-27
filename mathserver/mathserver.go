package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Starting Math Server......")
	http.HandleFunc("/squared", squaredHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func squaredHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called!")
	r.ParseForm()
	value := r.Form["value"]
	numvalue, _ := strconv.Atoi(value[0])
	fmt.Println("Squaring " + value[0])
	fmt.Fprintf(w, "%d", numvalue*numvalue)
}
