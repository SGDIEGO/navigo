package test

import (
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"
	"time"

	"github.com/SGDIEGO/Navigo/navigo"
)

func TestRouterSpeed(t *testing.T) {
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

	// Start a test HTTP server
	ts := httptest.NewServer(router1)
	defer ts.Close()

	timeReq := 2000
	// Send timeReq requests to random routes
	start := time.Now()
	for i := 0; i < timeReq; i++ {
		path := "/router2/router21"
		resp, err := http.Get(ts.URL + path)
		if err != nil {
			t.Fatalf("Error sending request: %v", err)
		}
		resp.Body.Close()
	}
	elapsed := time.Since(start)

	t.Logf("Elapsed time for %v requests: %s", timeReq, elapsed)
}
