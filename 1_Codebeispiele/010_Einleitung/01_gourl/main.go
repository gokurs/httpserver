package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	flagURL = flag.String("url", "", "the url to load")
)

func main() {
	flag.Parse()
	resp, err := http.Get(*flagURL)
	if err != nil {
		fmt.Printf("error loading %s\n%s", *flagURL, err)
	}

	responseOutput(resp)
	resp.Body.Close()

}

func responseOutput(resp *http.Response) {
	fmt.Fprintln(os.Stdout, "URL: ", *flagURL)
	fmt.Fprintln(os.Stdout, "Response:\n____________________")
	fmt.Fprintln(
		os.Stdout,
		formatResponseString(fmt.Sprintf("%#v", resp)),
	)
}

func formatResponseString(s string) string {
	s = strings.Replace(s, ", ", ",\n\n", -1)
	s = strings.Replace(s, ":", ":\n", -1)
	return s
}
