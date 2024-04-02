package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

func main() {
	fmt.Println("HELLO1004")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " World", time.Now())
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}