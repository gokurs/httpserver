package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Hier ist Text")
	printContent(s)
	f, err := os.Open("text.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("error when open file", err)
	}
	printContent(f)
	wp, err := http.Get("http://golang.org")
	if err != nil {
		log.Fatal("error loading webpage", err)
	}
	printContent(wp.Body)
	wp.Body.Close()

}

func printContent(r io.Reader) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("error reading", err)
	}
	fmt.Println(string(b))
}
