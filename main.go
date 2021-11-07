package main

import (
	"D:\coffee\handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// webserver on port 8080, using default handler
	http.ListenAndServe(":8000", sm)
}
