package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello world<h1>")
}

func main() {
	/**

	加两行注释
	*/

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)

}
