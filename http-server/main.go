package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Navaneethan_N")
	})

	http.HandleFunc("/Rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211CS224")
	})

	http.HandleFunc("/dept", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "COMPUTER SCIENCE AND ENGINEERING")
	})

	http.HandleFunc("/year", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "3rd")
	})

	http.HandleFunc("/address", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sathy - Bhavani State Highway, Alathukombai, Post, Sathyamangalam, Tamil Nadu 638401")
	})
 
	http.HandleFunc("/email", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "navaneethan.cs21@bitsathy.ac.in")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
