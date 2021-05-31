package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello world!")
		if err != nil {
			fmt.Println(err)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
