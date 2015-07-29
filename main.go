package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	host := os.Getenv("DEST_HOST")

	url, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        proxy,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Proxy to %s, Listening on port %s\n", host, port)
	log.Fatal(server.ListenAndServe())
}
