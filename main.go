package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Home Page!")
	if err != nil {
		fmt.Println(err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(1, 2)
	_, err := fmt.Fprintf(w, "About Page")
	fmt.Fprintf(w, fmt.Sprintf("Sum: %d", sum))
	if err != nil {
		fmt.Println(err)
	}
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Aplication open on port %s", port))

	_ = http.ListenAndServe(port, nil)
}
