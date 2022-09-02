package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func AddValues(x, y int) (int, error)  {
	var sum = x + y
	return sum, nil
}

func main() {


	// fmt.Println("Hello World!!!")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Hello World!!")
	// 	n, error := fmt.Fprintf(w, "Hello World!!!")
	// 	if error != nil {
	// 		fmt.Println(error)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Nuber of bytes written: %d", n))

	// })

	http.HandleFunc("/", Home)

	_ = http.ListenAndServe(":8080", nil)
}
