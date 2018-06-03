package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	srv := &http.Server{
		Addr:         ":1313",
		Handler:      nil,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
	}
	log.Fatal(srv.ListenAndServe())
}
