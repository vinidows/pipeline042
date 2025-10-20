package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Deploy realizado muito bem feito muito bonito")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
