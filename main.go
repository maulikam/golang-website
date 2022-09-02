package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, error := AddValues(2, 5)
	if error != nil {
		fmt.Println(error)
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, and add values %d", sum))
}

func Div(w http.ResponseWriter, r *http.Request) {

	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 10.00, 0.00, f))

}

func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("can't divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// AddValues adds two integers and return the sum
func AddValues(x, y int) (int, error) {
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
	http.HandleFunc("/About", About)
	http.HandleFunc("/Div", Div)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
