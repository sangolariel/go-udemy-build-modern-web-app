package main

import "fmt"

func main() {
	color := "Green"
	fmt.Println("Color before set: ", color)
	changeColor(&color)
	fmt.Println("Color after set: ", color)
}

func changeColor(s *string) {
	fmt.Println("Color before set: ", s)
	*s = "Red"
}
