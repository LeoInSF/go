package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1> Hey, there!</h1>
		<p> Go is fast!</p>
		<p> ... and simple!</p>
	`)

	fmt.Fprintf(w, "<h1> Hey, there!</h1>")
	fmt.Fprintf(w, "<p> Go is fast!</p>")
	fmt.Fprintf(w, "<p> ... and simple!</p>")
	fmt.Fprintf(w, "<p> You %s even add %s </p>", "can", "<strong> variables </strong>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Exper web designed by Leo_Ye")
}
