package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	b := &bytes.Buffer{}
	writeText(b)
	fmt.Println(b.String())
}

func writeText(w io.Writer) {
	_, err := fmt.Fprintln(w, "Ein wenig Text")
	if err != nil {
		log.Fatal("error writing text", err)
	}
}
