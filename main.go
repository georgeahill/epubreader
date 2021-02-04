package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func helloEndpoint(w http.ResponseWriter, r *http.Request) {
	// takes the information passed in the URL of the get request and puts it in a map (dictonary)
	name := r.URL.Query().Get("name")

	if name == "" {
		// return error 400 BAD REQUEST
		log.Println("400 name not found")
		http.Error(w, "name not found", http.StatusBadRequest)
		return
	}

	// return value to caller using "file print format"
	fmt.Fprintf(w, "Hello, %v!", name)
}

func addNumbers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	numbers, ok := query["number"]

	if !ok {
		log.Println("400 no numbers")
		http.Error(w, "Numbers not sent", http.StatusBadRequest)
		return
	}

	var result int
	for i, numTxt := range numbers {
		num, err := strconv.Atoi(numTxt)

		if err != nil {
			log.Println("400 " + strconv.Itoa(i) + " not number")
			http.Error(w, "Number "+strconv.Itoa(i)+" not number", http.StatusBadRequest)
			return
		}

		result += num
	}

	fmt.Fprintf(w, "Result: %v!\n", result)

}

// when you run remember & is a special symbol in bash add ""
// GET localhost:8080/add?number=1&number=4&...

func main() {
	fmt.Println("Starting server at port 8080")

	http.HandleFunc("/hello", helloEndpoint)
	http.HandleFunc("/add", addNumbers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
