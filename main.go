package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/niklastomas/go-product-api/handlers"
)

func main() {

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// Create handlers
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create new server
	s := http.Server{
		Addr:         ":8080",
		Handler:      ph,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	s.ListenAndServe()

	// go func() {
	// 	l.Println("Starting server on port 8080")

	// 	l.Fatal(s.ListenAndServe())

	// }()

	// trap sigterm or interrupt and gracefully shutdown the server
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)

}
