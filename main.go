package main

import (
	"errors"
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

func Devide(w http.ResponseWriter, r *http.Request) {
	f, err := devideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Can not divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f devided %f is % f", 100.0, 10.0, f))
}

func devideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can not devide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", Devide)

	fmt.Println(fmt.Sprintf("Aplication open on port %s", port))

	_ = http.ListenAndServe(port, nil)
}
