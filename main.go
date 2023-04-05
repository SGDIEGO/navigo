package main

import (
	"net/http"

	"github.com/SGDIEGO/Navigo/navigo"
)

func main() {
	router1 := navigo.NewMux()

	router1.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from " + r.URL.Path))
	})

	router1.GET("/U", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from " + r.URL.Path))
	})

	router2 := router1.Group("/router2")
	router2.GET("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from " + r.URL.Path))
	})

	router2_1 := router2.Group("/router21")
	router2_1.GET("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from " + r.URL.Path))
	})

	router3 := router1.Group("/router3")
	router3.GET("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from " + r.URL.Path))
	})

	http.ListenAndServe(":3000", router1)
}
