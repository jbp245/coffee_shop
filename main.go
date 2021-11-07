package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Bad Request", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Response Body: %s\n", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	// webserver on port 8080, using default handler
	http.ListenAndServe(":8000", nil)

}
